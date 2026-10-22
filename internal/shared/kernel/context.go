package kernel

import "context"

type ctxKey string

const actorKey ctxKey = "actor"

// WithActor stores the actor in context.
func WithActor(ctx context.Context, a *Actor) context.Context {
	return context.WithValue(ctx, actorKey, a)
}

// ActorFromContext retrieves the actor from context.
func ActorFromContext(ctx context.Context) (*Actor, bool) {
	a, ok := ctx.Value(actorKey).(*Actor)
	return a, ok
}

// MustActor retrieves the actor or panics if missing.
func MustActor(ctx context.Context) *Actor {
	a, ok := ActorFromContext(ctx)
	if !ok {
		panic("missing actor in context")
	}
	return a
}
