#EML(RFC5322) PARSER

*This aims to be a more complete implementation of go's net/mail package.*


This is parser of message in rfc5322(.eml)

Split message on all headers(subject,from,to,cc,date,et.c.), body text, body html, and attachments(decodes it)
