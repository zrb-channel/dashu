package dashu

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/zrb-channel/core/schema"
	"github.com/zrb-channel/utils"
	log "github.com/zrb-channel/utils/logger"
)

func dataIn(ctx context.Context, conf *Config, data string) (*DataInResult, error) {
	resp, err := utils.Request(ctx).SetBody(&DataInRequest{
		PublicKey:  conf.DsPublicKey,
		PrivateKey: conf.PrivateKey,
		Data:       data,
	}).Post(conf.ServiceAddr + "/in")

	if err != nil {
		return nil, err
	}

	res := &schema.ServiceBaseResponse[*DataInResult]{}

	if err = json.Unmarshal(resp.Body(), res); err != nil {
		return nil, err
	}

	return res.Data, nil
}

func dataOut(ctx context.Context, conf *Config, data, sign string) ([]byte, error) {

	resp, err := utils.Request(ctx).SetBody(&DataOutRequest{
		PublicKey:  conf.DsPublicKey,
		PrivateKey: conf.PrivateKey,
		Data:       data,
		Sign:       sign,
	}).Post(conf.ServiceAddr + "/out")

	if err != nil {
		return nil, err
	}

	res := &schema.ServiceBaseResponse[json.RawMessage]{}
	if err = json.Unmarshal(resp.Body(), res); err != nil {
		log.WithError(err).Error("数据解析失败", log.Fields(map[string]any{"resp": resp.String()}))
		return nil, err
	}

	if res.Code != 200 {
		return nil, errors.New(res.Msg)
	}
	return res.Data, nil
}

func Apply(ctx context.Context, conf *Config, orderNo string) (*ApplyResponse, error) {

	req := &ApplyRequest{
		ChannelId: conf.ChannelId,
		ProductId: conf.ProductId,
		OrderNo:   orderNo,
	}

	body := &BaseRequest{
		ChannelId: conf.ChannelId,
	}

	msg, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	dataInRes, err := dataIn(ctx, conf, string(msg))
	if err != nil {
		return nil, err
	}
	body.Data = dataInRes.Encrypt
	body.Sign = dataInRes.Sign

	resp, err := utils.Request(ctx).SetBody(body).Post(applyAddr)
	if err != nil {
		log.WithError(err).Error("请求失败", log.Fields(map[string]string{"addr": applyAddr}))
		return nil, err
	}

	res := &BaseResponse{}
	if err = json.Unmarshal(resp.Body(), res); err != nil {
		return nil, err
	}

	result, err := dataOut(ctx, conf, res.Data, res.Sign)
	if err != nil {
		return nil, err
	}

	data := &ApplyResponse{}
	if err = json.Unmarshal(result, data); err != nil {
		return nil, err
	}

	return data, nil
}
