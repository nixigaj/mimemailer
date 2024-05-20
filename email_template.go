package mimemailer

const emailTemplate = `Subject: {{ .Subject }}
From: {{ .From }}
To: {{ .To }}
Date: {{ .Date }}{{ if .ListUnsubscribe }}
List-Unsubscribe: {{ .ListUnsubscribe }}{{ end }}
MIME-Version: 1.0
Content-Type: multipart/alternative; boundary=boundary42

--boundary42
Content-Type: text/plain; charset=utf-8
Content-Transfer-Encoding: quoted-printable

{{ .TextQP }}

--boundary42
Content-Type: text/html; charset=utf-8
Content-Transfer-Encoding: quoted-printable

{{ .HTMLQP }}

--boundary42--
`
