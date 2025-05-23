package handler

import (
	"net/http"

	"shortLink/api/internal/logic"
	"shortLink/api/internal/svc"
	"shortLink/api/internal/types"

	"github.com/go-playground/validator/v10"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func SearchHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SearchRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		// 处理参数
		validate := validator.New()
		if err := validate.StructCtx(r.Context(), &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewSearchLogic(r.Context(), svcCtx)
		resp, err := l.Search(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			// 跳转
			longURL := resp.LongURL
			http.Redirect(w, r, longURL, http.StatusFound)
		}
	}
}
