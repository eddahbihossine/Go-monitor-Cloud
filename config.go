package main

type Config struct {
    CPUThreshold    float64
    MemoryThreshold float64
    DiskThreshold   float64
    EmailConfig     EmailSettings
    Mock            bool
}

type EmailSettings struct {
    From     string
    To       string
    Password string
    SMTPHost string
    SMTPPort int
}

func GetDefaultConfig() Config {
    return Config{
        CPUThreshold:    80.0,
        MemoryThreshold: 80.0,
        DiskThreshold:   90.0,
        Mock:            false,
        EmailConfig: EmailSettings{
            From:     "your-email@gmail.com",
            To:       "destination-email@example.com",
            Password: "your-app-password", // App password or real one (Gmail: use app password)
            SMTPHost: "smtp.gmail.com",
            SMTPPort: 587,
        },
    }
}
