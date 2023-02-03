package controllers

import (
	"_template_/domain/repository"
	"_template_/registry"
	"context"
	"net/http"
)

type TransferReq struct {
	OutID  string  `json:"out_id"`
	InID   string  `json:"in_id"`
	Amount float64 `json:"amount"`
}

func Transfer(ctx context.Context, req TransferReq) (string, int, error) {
	out, err := registry.Repository().Users().FindOne(ctx, req.OutID)
	if err != nil {
		return "", http.StatusNotFound, err
	}

	in, err := registry.Repository().Users().FindOne(ctx, req.InID)
	if err != nil {
		return "", http.StatusNotFound, err
	}

	err = registry.UOW().Atomic(ctx, func(r repository.Repository) error {
		in.AddAmount(req.Amount)
		out.ReduceAmount(req.Amount)
		if err := r.Users().Save(ctx, &in); err != nil {
			return err
		}
		if err := r.Users().Save(ctx, &out); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return "", http.StatusInternalServerError, err
	}

	return "success", http.StatusOK, nil
}
