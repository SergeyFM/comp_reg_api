package main

import "net/http"

func RegisterSwaggerRoutes(mux *http.ServeMux) {

	mux.Handle("/swagger-ui/",
		http.StripPrefix("/swagger-ui/",
			http.FileServer(http.Dir("./swagger-ui")),
		),
	)

	mux.HandleFunc("GET /openapi.yaml", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "openapi.yaml")
	})

	mux.HandleFunc("GET /swagger", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")

		w.Write([]byte(`<!doctype html>
<html>
<head>
  <title>Company Registry API Swagger</title>
  <link rel="stylesheet" href="/swagger-ui/swagger-ui.css">
</head>
<body>
  <div id="swagger-ui"></div>

  <script src="/swagger-ui/swagger-ui-bundle.js"></script>

  <script>
    window.onload = function() {
      SwaggerUIBundle({
        url: '/openapi.yaml',
        dom_id: '#swagger-ui'
      });
    };
  </script>
</body>
</html>`))
	})
}
