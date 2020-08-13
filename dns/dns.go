package dns

import (
	"fmt"
	"time"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/errors"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/alidns"
	"github.com/quteam/aliddns/logger"
	"github.com/quteam/aliddns/utils"
)

type Handler struct {
	client *alidns.Client
	ip     string
	domain string
	rr     string
}

func New(domain, ip, rr string) *Handler {
	if domain == "" || ip == "" || rr == "" {
		panic(fmt.Errorf("domain, ip or rr cannot be empty"))
	}
	client, err := alidns.NewClientWithAccessKey(utils.Env("REGION", "cn-hangzhou"), utils.Env("ACCESS_KEY"), utils.Env("ACCESS_KEY_SECRET"))
	if err != nil {
		panic(fmt.Errorf("new alidns client failed: %v", err))
	}
	return &Handler{
		client: client,
		ip:     ip,
		domain: domain,
		rr:     rr,
	}
}

func (dns *Handler) findRecords() (*alidns.DescribeDomainRecordsResponse, error) {
	reqest := alidns.CreateDescribeDomainRecordsRequest()
	reqest.DomainName = dns.domain
	resp, err := dns.client.DescribeDomainRecords(reqest)
	if err != nil {
		// try to fix timeout issue
		if clientErr, ok := err.(*errors.ClientError); ok && clientErr.ErrorCode() == errors.TimeoutErrorCode {
			// retry
			logger.Error("timeout. retry...", clientErr)
			time.Sleep(time.Second)
			return dns.findRecords()
		}
		logger.Error("finding records failed", err)
		return nil, fmt.Errorf("finding records failed: %v", err)
	}
	return resp, nil
}

func (dns *Handler) addRecord() (*alidns.AddDomainRecordResponse, error) {
	request := alidns.CreateAddDomainRecordRequest()
	request.DomainName = dns.domain
	request.Type = "A"
	request.RR = dns.rr
	request.Value = dns.ip
	resp, err := dns.client.AddDomainRecord(request)
	if err != nil {
		logger.Error("adding record failed", err)
		return nil, fmt.Errorf("adding record failed: %v", err)
	}
	logger.Info(fmt.Sprintf(`set ip of '%s.%s' to %s`, dns.rr, dns.domain, dns.ip))
	return resp, nil
}

func (dns *Handler) updateRecord(recordId string) (*alidns.UpdateDomainRecordResponse, error) {
	request := alidns.CreateUpdateDomainRecordRequest()
	request.RecordId = recordId
	request.Type = "A"
	request.RR = dns.rr
	request.Value = dns.ip
	resp, err := dns.client.UpdateDomainRecord(request)
	if err != nil {
		logger.Error("updating record failed", err)
		return nil, fmt.Errorf("updating record failed: %v", err)
	}
	logger.Info(fmt.Sprintf(`set ip of '%s.%s' to %s`, dns.rr, dns.domain, dns.ip))
	return resp, nil
}

func (dns *Handler) Bind() error {
	logger.Info(fmt.Sprintf("current ip is %s", dns.ip))
	recordResp, err := dns.findRecords()
	if err != nil {
		return err
	}
	records := recordResp.DomainRecords.Record
	shouldAdd := true
	var recordId, recordValue string
	for _, r := range records {
		if r.RR == dns.rr {
			// 如果找到RR和输入里的rr相同的记录，则更新这条记录的解析。反之则添加一条新解析
			shouldAdd = false
			recordId = r.RecordId
			recordValue = r.Value
			break
		}
	}
	// add
	if shouldAdd {
		logger.Info("add domain record")
		if _, err := dns.addRecord(); err != nil {
			return err
		}
		return nil
	}
	// update record
	logger.Info(fmt.Sprintf("domain ip is %s", recordValue))
	if recordValue == dns.ip {
		// no need updating
		logger.Info("ip not changed, no need updating")
		return nil
	}
	logger.Info("ip changed, update domain record")
	if _, err := dns.updateRecord(recordId); err != nil {
		return err
	}
	return nil
}