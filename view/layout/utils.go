package layout

import "context"

func GetUserFromContext(ctx context.Context) string {
	user, ok := ctx.Value("user").(string)

	if !ok {
		return ""
	}

	return user
}
