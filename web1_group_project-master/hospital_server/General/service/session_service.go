package service

import (
	"github.com/web1_group_project/hospital_server/General"
	"github.com/web1_group_project/hospital_server/entity"
)

// SessionServiceImpl implements user.SessionService interface
type SessionServiceImpl struct {
	sessionRepo General.SessionRepository
}

// NewSessionService  returns a new SessionService object
func NewSessionService(sessRepository General.SessionRepository) General.SessionService {
	return &SessionServiceImpl{sessionRepo: sessRepository}
}

// Session returns a given stored session
func (ss *SessionServiceImpl) Session(sessionID string) (*entity.Session, []error) {
	sess, errs := ss.sessionRepo.Session(sessionID)
	if len(errs) > 0 {
		return nil, errs
	}
	return sess, errs
}

// StoreSession stores a given session
func (ss *SessionServiceImpl) StoreSession(session *entity.Session) (*entity.Session, []error) {
	sess, errs := ss.sessionRepo.StoreSession(session)
	if len(errs) > 0 {
		return nil, errs
	}
	return sess, errs
}

// DeleteSession deletes a given session
func (ss *SessionServiceImpl) DeleteSession(sessionID string) (*entity.Session, []error) {
	sess, errs := ss.sessionRepo.DeleteSession(sessionID)
	if len(errs) > 0 {
		return nil, errs
	}
	return sess, errs
}
