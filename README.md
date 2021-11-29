![Test](https://img.shields.io/static/v1?message=success&color=%3Cgreen%3E)
![Issues](https://img.shields.io/github/issues/iradukunda1/email-client.svg)
![Lint](https://img.shields.io/static/v1?message=success&color=%3Cgreen%3E)
[![codecov](https://codecov.io/gh/iradukunda1/email-client/branch/master/graph/badge.svg)](https://codecov.io/gh/iradukunda1/email-client)

# email-client

The email-client Implement send email based on gmail protocol

### Usage
Run the following command to install the package
```
go get github.com/iradukunda1/email-client

```
```bash
		config := &config{
			Host:   "smtp.gmail.com",
			Port:   "587",
			Sender: "your-email@gmail.com",
			Secret: "******************" // your gmail password,
		}
        // descructuring our config you dont pass it directly
		configs := email.Config{
			Host:   config.Host,
			Port:   config.Port,
			Sender: config.Sender,
			Secret: config.Secret,
		}

		sms := &message{
			Subject:    "testing",
			Body:       "Hello this is testing message"
			Sender:     config.Sender,
			Recipients: []string{"receiver@gmail.com"},
		}

		message := email.Email{
			Subject:    sms.Subject,
			Body:       sms.Body,
			Sender:     sms.Sender,
			Recipients: sms.Recipients,
		}

		err = email.Send(&configs, &message)

		if err != nil {
			return err
		}
```
