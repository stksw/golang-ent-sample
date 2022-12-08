// Code generated by ent, DO NOT EDIT.

package http

import (
	"ent-sample/ent"
	"ent-sample/ent/todo"
	"ent-sample/ent/user"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/mailru/easyjson"
	"go.uber.org/zap"
)

// Update updates a given ent.Todo and saves the changes to the database.
func (h TodoHandler) Update(w http.ResponseWriter, r *http.Request) {
	l := h.log.With(zap.String("method", "Update"))
	// ID is URL parameter.
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		l.Error("error getting id from url parameter", zap.String("id", chi.URLParam(r, "id")), zap.Error(err))
		BadRequest(w, "id must be an integer")
		return
	}
	// Get the post data.
	var d TodoUpdateRequest
	if err := easyjson.UnmarshalFromReader(r.Body, &d); err != nil {
		l.Error("error decoding json", zap.Error(err))
		BadRequest(w, "invalid json string")
		return
	}
	// Validate the data.
	errs := make(map[string]string)
	if d.Text == nil {
		errs["text"] = `missing required field: "text"`
	} else if err := todo.TextValidator(*d.Text); err != nil {
		errs["text"] = strings.TrimPrefix(err.Error(), "todo: ")
	}
	if d.Status != nil {
		if err := todo.StatusValidator(*d.Status); err != nil {
			errs["status"] = strings.TrimPrefix(err.Error(), "todo: ")
		}
	}
	if len(errs) > 0 {
		l.Info("validation failed", zapFields(errs)...)
		BadRequest(w, errs)
		return
	}
	// Save the data.
	b := h.client.Todo.UpdateOneID(id)
	if d.Text != nil {
		b.SetText(*d.Text)
	}
	if d.Status != nil {
		b.SetStatus(*d.Status)
	}
	if d.Priority != nil {
		b.SetPriority(*d.Priority)
	}
	if d.Children != nil {
		b.ClearChildren().AddChildIDs(d.Children...)
	}
	if d.Parent != nil {
		b.SetParentID(*d.Parent)

	}
	// Store in database.
	e, err := b.Save(r.Context())
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			msg := stripEntError(err)
			l.Info(msg, zap.Error(err), zap.Int("id", id))
			NotFound(w, msg)
		case ent.IsNotSingular(err):
			msg := stripEntError(err)
			l.Error(msg, zap.Error(err), zap.Int("id", id))
			BadRequest(w, msg)
		default:
			l.Error("could-not-update-todo", zap.Error(err), zap.Int("id", id))
			InternalServerError(w, nil)
		}
		return
	}
	// Reload entry.
	q := h.client.Todo.Query().Where(todo.ID(e.ID))
	e, err = q.Only(r.Context())
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			msg := stripEntError(err)
			l.Info(msg, zap.Error(err), zap.Int("id", id))
			NotFound(w, msg)
		case ent.IsNotSingular(err):
			msg := stripEntError(err)
			l.Error(msg, zap.Error(err), zap.Int("id", id))
			BadRequest(w, msg)
		default:
			l.Error("could-not-read-todo", zap.Error(err), zap.Int("id", id))
			InternalServerError(w, nil)
		}
		return
	}
	l.Info("todo rendered", zap.Int("id", id))
	easyjson.MarshalToHTTPResponseWriter(NewTodo1977301413View(e), w)
}

// Update updates a given ent.User and saves the changes to the database.
func (h UserHandler) Update(w http.ResponseWriter, r *http.Request) {
	l := h.log.With(zap.String("method", "Update"))
	// ID is URL parameter.
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		l.Error("error getting id from url parameter", zap.String("id", chi.URLParam(r, "id")), zap.Error(err))
		BadRequest(w, "id must be an integer")
		return
	}
	// Get the post data.
	var d UserUpdateRequest
	if err := easyjson.UnmarshalFromReader(r.Body, &d); err != nil {
		l.Error("error decoding json", zap.Error(err))
		BadRequest(w, "invalid json string")
		return
	}
	// Validate the data.
	errs := make(map[string]string)
	if d.Name == nil {
		errs["name"] = `missing required field: "name"`
	} else if err := user.NameValidator(*d.Name); err != nil {
		errs["name"] = strings.TrimPrefix(err.Error(), "user: ")
	}
	if d.Email == nil {
		errs["email"] = `missing required field: "email"`
	} else if err := user.EmailValidator(*d.Email); err != nil {
		errs["email"] = strings.TrimPrefix(err.Error(), "user: ")
	}
	if d.Password == nil {
		errs["password"] = `missing required field: "password"`
	} else if err := user.PasswordValidator(*d.Password); err != nil {
		errs["password"] = strings.TrimPrefix(err.Error(), "user: ")
	}
	if d.Status != nil {
		if err := user.StatusValidator(*d.Status); err != nil {
			errs["status"] = strings.TrimPrefix(err.Error(), "user: ")
		}
	}
	if len(errs) > 0 {
		l.Info("validation failed", zapFields(errs)...)
		BadRequest(w, errs)
		return
	}
	// Save the data.
	b := h.client.User.UpdateOneID(id)
	if d.Name != nil {
		b.SetName(*d.Name)
	}
	if d.Email != nil {
		b.SetEmail(*d.Email)
	}
	if d.Password != nil {
		b.SetPassword(*d.Password)
	}
	if d.Status != nil {
		b.SetStatus(*d.Status)
	}
	// Store in database.
	e, err := b.Save(r.Context())
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			msg := stripEntError(err)
			l.Info(msg, zap.Error(err), zap.Int("id", id))
			NotFound(w, msg)
		case ent.IsNotSingular(err):
			msg := stripEntError(err)
			l.Error(msg, zap.Error(err), zap.Int("id", id))
			BadRequest(w, msg)
		default:
			l.Error("could-not-update-user", zap.Error(err), zap.Int("id", id))
			InternalServerError(w, nil)
		}
		return
	}
	// Reload entry.
	q := h.client.User.Query().Where(user.ID(e.ID))
	e, err = q.Only(r.Context())
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			msg := stripEntError(err)
			l.Info(msg, zap.Error(err), zap.Int("id", id))
			NotFound(w, msg)
		case ent.IsNotSingular(err):
			msg := stripEntError(err)
			l.Error(msg, zap.Error(err), zap.Int("id", id))
			BadRequest(w, msg)
		default:
			l.Error("could-not-read-user", zap.Error(err), zap.Int("id", id))
			InternalServerError(w, nil)
		}
		return
	}
	l.Info("user rendered", zap.Int("id", id))
	easyjson.MarshalToHTTPResponseWriter(NewUser2774811467View(e), w)
}