// Code generated by ogen, DO NOT EDIT.

package api

import (
	"context"

	ht "github.com/ogen-go/ogen/http"
)

// UnimplementedHandler is no-op Handler which returns http.ErrNotImplemented.
type UnimplementedHandler struct{}

var _ Handler = UnimplementedHandler{}

// CreateAnnotation implements createAnnotation operation.
//
// Create a new annotation.
//
// POST /annotations
func (UnimplementedHandler) CreateAnnotation(ctx context.Context, req *NewAnnotation) (r CreateAnnotationRes, _ error) {
	return r, ht.ErrNotImplemented
}

// DeleteAnnotation implements deleteAnnotation operation.
//
// Delete an annotation by id.
//
// DELETE /annotations/{id}
func (UnimplementedHandler) DeleteAnnotation(ctx context.Context, params DeleteAnnotationParams) (r DeleteAnnotationRes, _ error) {
	return r, ht.ErrNotImplemented
}

// GetAnnotation implements getAnnotation operation.
//
// Get a single annotation by id.
//
// GET /annotations/{id}
func (UnimplementedHandler) GetAnnotation(ctx context.Context, params GetAnnotationParams) (r GetAnnotationRes, _ error) {
	return r, ht.ErrNotImplemented
}

// GetAnnotations implements getAnnotations operation.
//
// Get all annotations and filter by datasets or timerange.
//
// GET /annotations
func (UnimplementedHandler) GetAnnotations(ctx context.Context, params GetAnnotationsParams) (r GetAnnotationsRes, _ error) {
	return r, ht.ErrNotImplemented
}

// UpdateAnnotation implements updateAnnotation operation.
//
// Update an annotation.
//
// PUT /annotations/{id}
func (UnimplementedHandler) UpdateAnnotation(ctx context.Context, req *UpdateAnnotation, params UpdateAnnotationParams) (r UpdateAnnotationRes, _ error) {
	return r, ht.ErrNotImplemented
}
