package proton

import (
	"context"

	"github.com/go-resty/resty/v2"
)

func (c *Client) ListChildren(ctx context.Context, shareID, linkID string) ([]Link, error) {
	var res struct {
		Links []Link
	}

	if err := c.do(ctx, func(r *resty.Request) (*resty.Response, error) {
		return r.SetResult(&res).Get("/drive/shares/" + shareID + "/folders/" + linkID + "/children")
	}); err != nil {
		return nil, err
	}

	return res.Links, nil
}

func (c *Client) CreateFolder(ctx context.Context, shareID string, req CreateFolderReq) (CreateFolderRes, error) {
	var res struct {
		File CreateFolderRes
	}

	if err := c.do(ctx, func(r *resty.Request) (*resty.Response, error) {
		return r.SetResult(&res).SetBody(req).Post("/drive/shares/" + shareID + "/folders")
	}); err != nil {
		return CreateFolderRes{}, err
	}

	return res.File, nil
}