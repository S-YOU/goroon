package goroon

import (
	"os"
	"testing"
	"time"

	"github.com/tzmfreedom/goroon"

	"gopkg.in/h2non/gock.v1"
)

func TestGetScheduleByUserId(t *testing.T) {
	defer gock.Off()

	gock.New("https://garoon.com").
		Post("/cbpapi/schedule/api").
		Reply(200).BodyString(`<?xml version="1.0" encoding="utf-8"?>
<soap:Envelope
xmlns:soap="http://www.w3.org/2003/05/soap-envelope"
xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
xmlns:xsd="http://www.w3.org/2001/XMLSchema"
xmlns:schedule="http://wsdl.cybozu.co.jp/schedule/2008">
<soap:Header>
    <vendor>Cybozu</vendor>
    <product>Garoon</product>
    <product_type>1</product_type>
    <version>3.7.5</version>
    <apiversion>1.3.1</apiversion>
</soap:Header>
<soap:Body>
    <schedule:ScheduleGetEventsByTargetResponse>
        <returns>
            <schedule_event id="123"
                detail="fugafuga"
                description="hogehoge"
                >
                <members xmlns="http://schemas.cybozu.co.jp/schedule/2008">
                    <member>
                        <user id="aa" name="bb" order="0" />
                    </member>
                </members>
                <repeat_info xmlns="http://schemas.cybozu.co.jp/schedule/2008">
                <condition type="week" day="20"
                    week="2" start_date="2016-11-22" end_date="2017-04-01"
                    start_time="14:00:00" end_time="14:30:00"/>
                    <exclusive_datetimes>
                        <exclusive_datetime start="2016-12-13T00:00:00+09:00" end="2016-12-14T00:00:00+09:00" />
                        <exclusive_datetime start="2016-12-20T00:00:00+09:00" end="2016-12-21T00:00:00+09:00" />
                    </exclusive_datetimes>
                </repeat_info>
                <when>
                    <datetime start="2016-12-15T13:07:00Z" end="2016-12-15T16:30:00Z" />
                </when>
            </schedule_event>
        </returns>
    </schedule:ScheduleGetEventsByTargetResponse>
</soap:Body>
</soap:Envelope>`)

	client := goroon.NewClient("username", "password", "https://garoon.com", true, os.Stdout)
	res := &goroon.ScheduleGetEventsByTargetResponse{}

	tm := time.Now()
	req := &goroon.ScheduleGetEventsByTargetRequest{
		Parameters: &goroon.Parameters{
			Start: &tm,
			End:   &tm,
			User: &goroon.User{
				Id: "userId",
			},
		},
	}
	client.ScheduleGetEventsByTarget(req, res)
}