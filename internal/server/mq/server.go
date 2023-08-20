package server

type Server struct {
	auditServer *AuditServer
}

func NewServer(auditServer *AuditServer) *Server {
	return &Server{
		auditServer: auditServer,
	}
}

func (s *Server) ListenAndServe() error {
	var err error

	go func() {
		err = s.auditServer.ListenAndServe()
	}()
	if err != nil {
		return err
	}

	return nil
}
