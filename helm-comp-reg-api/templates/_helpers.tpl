{{- define "comp-reg-api.name" -}}
comp-reg-api
{{- end }}

{{- define "comp-reg-api.fullname" -}}
{{ .Release.Name }}-comp-reg-api
{{- end }}

{{- define "comp-reg-api.labels" -}}
app.kubernetes.io/name: {{ include "comp-reg-api.name" . }}
app.kubernetes.io/instance: {{ .Release.Name }}
app.kubernetes.io/version: {{ .Chart.AppVersion }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
{{- end }}

{{- define "comp-reg-api.selectorLabels" -}}
app.kubernetes.io/name: {{ include "comp-reg-api.name" . }}
app.kubernetes.io/instance: {{ .Release.Name }}
{{- end }}