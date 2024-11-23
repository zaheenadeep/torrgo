package tracker

import (
	"net/http"
	"github.com/zaheenandeep/torrgo/internal/metainfo"
)

func Initiate(mi *metainfo.Metainfo) {
	c := &http.Client{}
	req := c.NewRequest(mi.)
}
