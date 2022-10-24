package dashu

type Config struct {
	ServiceAddr string `json:"serviceAddr"`
	ChannelId   string `json:"channelId"`
	PublicKey   string `json:"publicKey"`
	PrivateKey  string `json:"privateKey"`
	DsPublicKey string `json:"DsPublicKey"`
	ProductId   string `json:"productId"`
}

type (
	DataInRequest struct {
		PublicKey  string `json:"publicKey"`
		PrivateKey string `json:"privateKey"`
		Data       string `json:"data"`
	}

	DataOutRequest struct {
		PublicKey  string `json:"publicKey"`
		PrivateKey string `json:"privateKey"`
		Data       string `json:"data"`
		Sign       string `json:"sign"`
	}

	DataInResult struct {
		Sign    string `json:"sign"`
		Encrypt string `json:"encrypt"`
	}
)

type BaseRequest struct {
	ChannelId string `json:"channelId"`
	Data      string `json:"data"`
	Sign      string `json:"sign"`
}

type BaseResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    string `json:"data"`
	Sign    string `json:"sign"`
}

type ApplyRequest struct {
	ChannelId string `json:"channelId"`
	ProductId string `json:"productId"`
	OrderNo   string `json:"outOrderId"`
}

type ApplyResponse struct {
	OrderNo     string `json:"outOrderId"`
	RedirectUrl string `json:"redirectUrl"`
	ChannelId   string `json:"channelId"`
}
