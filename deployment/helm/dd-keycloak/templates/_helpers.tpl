{{/*
Expand the name of the chart.
*/}}
{{- define "dd-keycloak.name" -}}
{{- default .Chart.Name .Values.nameOverride | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Create a default fully qualified app name.
We truncate at 63 chars because some Kubernetes name fields are limited to this (by the DNS naming spec).
If release name contains chart name it will be used as a full name.
*/}}
{{- define "dd-keycloak.fullname" -}}
{{- if .Values.fullnameOverride }}
{{- .Values.fullnameOverride | trunc 63 | trimSuffix "-" }}
{{- else }}
{{- $name := default .Chart.Name .Values.nameOverride }}
{{- if contains $name .Release.Name }}
{{- .Release.Name | trunc 63 | trimSuffix "-" }}
{{- else }}
{{- printf "%s-%s" .Release.Name $name | trunc 63 | trimSuffix "-" }}
{{- end }}
{{- end }}
{{- end }}

{{/*
Create chart name and version as used by the chart label.
*/}}
{{- define "dd-keycloak.chart" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Common labels
*/}}
{{- define "dd-keycloak.labels" -}}
helm.sh/chart: {{ include "dd-keycloak.chart" . }}
{{ include "dd-keycloak.selectorLabels" . }}
{{- if .Chart.AppVersion }}
app.kubernetes.io/version: {{ .Chart.AppVersion | quote }}
{{- end }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
{{- end }}

{{/*
Selector labels
*/}}
{{- define "dd-keycloak.selectorLabels" -}}
app.kubernetes.io/name: {{ include "dd-keycloak.name" . }}
app.kubernetes.io/instance: {{ .Release.Name }}
{{- end }}

{{/*
Create the name of the service account to use
*/}}
{{- define "dd-keycloak.serviceAccountName" -}}
{{- if .Values.serviceAccount.create }}
{{- default (include "dd-keycloak.fullname" .) .Values.serviceAccount.name }}
{{- else }}
{{- default "default" .Values.serviceAccount.name }}
{{- end }}
{{- end }}

{{- define "dd-keycloak.authApp.service.name" -}}
{{- "auth-app-svc" }}
{{- end -}}

{{- define "dd-keycloak.authApp.service.port" -}}
{{- "8081" }}
{{- end -}}

{{- define "dd-keycloak.nginx-rp.service.name" -}}
{{- "nginx-rp-svc" }}
{{- end -}}

{{- define "dd-keycloak.nginx-rp.service" -}}
{{- printf "http://%s:%v" (include "keycloak.fullname" .Subcharts.keycloak) .Values.keycloak.service.ports.http }}
{{- end -}}

{{- define "dd-keycloak.keycloak.service" -}}
{{- printf "http://%s:%v" (include "keycloak.fullname" .Subcharts.keycloak) .Values.keycloak.service.ports.http }}
{{- end -}}
