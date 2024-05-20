package mimemailer

const emailTemplate = `Subject: {{ .Subject }}
From: {{ .From }}
To: {{ .To }}
Reply-To: {{ .Name }} <{{ .ReplyTo }}>
Date: {{ .Date }}{{ if .ListUnsubscribe }}
List-Unsubscribe: {{ .ListUnsubscribe }}{{ end }}
MIME-Version: 1.0
Content-Type: multipart/alternative; boundary=multipart_email_boundary

--multipart_email_boundary
Content-Type: text/plain; charset=utf-8
Content-Transfer-Encoding: quoted-printable

{{ .TextQP }}

--multipart_email_boundary
Content-Type: text/html; charset=utf-8
Content-Transfer-Encoding: quoted-printable

{{ .HTMLQP }}

--multipart_email_boundary--
`
