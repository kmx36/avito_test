package handlers

import (
    "context"
)

func WithUserID(ctx context.Context, userID int) context.Context {
    return context.WithValue(ctx, "userID", userID)
}