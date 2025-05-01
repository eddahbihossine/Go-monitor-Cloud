package main

import (
	"fmt"
	"log"
	"net/smtp"
	"os"
	"time"
	"crypto/tls"

	"github.com/joho/godotenv"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/mem"
)

func checkMetrics() (float64, float64, float64, error) {
	cpuPercent, err := cpu.Percent(0, false)
	if err != nil {
		return 0, 0, 0, err
	}

	memStat, err := mem.VirtualMemory()
	if err != nil {
		return 0, 0, 0, err
	}

	diskStat, err := disk.Usage("/")
	if err != nil {
		return 0, 0, 0, err
	}

	return cpuPercent[0], memStat.UsedPercent, diskStat.UsedPercent, nil
}

// func sendEmail(subject, body string) error {
//     from := os.Getenv("EMAIL_FROM")
//     pass := os.Getenv("EMAIL_PASS")
//     to := os.Getenv("EMAIL_TO")
//     host := os.Getenv("EMAIL_HOST")
//     port := "587"

//     msg := fmt.Sprintf(
//         "From: %s\r\nTo: %s\r\nSubject: %s\r\n\r\n%s",
//         from, to, subject, body,
//     )

//     c, err := smtp.Dial(host + ":" + port)
//     if err != nil {
//         return fmt.Errorf("SMTP dial failed: %v", err)
//     }
//     defer c.Close()

//     if ok, _ := c.Extension("STARTTLS"); ok {
//         tlsConfig := &tls.Config{ServerName: host}
//         if err = c.StartTLS(tlsConfig); err != nil {
//             return fmt.Errorf("STARTTLS failed: %v", err)
//         }
//     } else {
//         return fmt.Errorf("STARTTLS not supported")
//     }

//     auth := smtp.PlainAuth("", from, pass, host)
//     if err = c.Auth(auth); err != nil {
//         return fmt.Errorf("auth failed: %v", err)
//     }

//     if err = c.Mail(from); err != nil {
//         return fmt.Errorf("MAIL command failed: %v", err)
//     }
//     if err = c.Rcpt(to); err != nil {
//         return fmt.Errorf("RCPT command failed: %v", err)
//     }

//     w, err := c.Data()
//     if err != nil {
//         return fmt.Errorf("DATA command failed: %v", err)
//     }
//     if _, err = w.Write([]byte(msg)); err != nil {
//         return fmt.Errorf("writing message failed: %v", err)
//     }
//     if err = w.Close(); err != nil {
//         return fmt.Errorf("closing writer failed: %v", err)
//     }

//     return nil
// }

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("⚠️ No .env file found, continuing with system env vars...")
	}

	for {
		cpuVal, memVal, diskVal, err := checkMetrics()
		if err != nil {
			log.Println("Error collecting metrics:", err)
			continue
		}

		log.Printf("CPU: %.2f%%, Memory: %.2f%%, Disk: %.2f%%\n", cpuVal, memVal, diskVal)

		// Example thresholds
		if cpuVal > 80 || memVal > 80 || diskVal > 90 {
			subject := "⚠️ Server Alert: High Resource Usage"
			body := fmt.Sprintf("CPU: %.2f%%\nMemory: %.2f%%\nDisk: %.2f%%", cpuVal, memVal, diskVal)
			err := sendEmail(subject, body)
			if err != nil {
				log.Println("❌ Failed to send email:", err)
			} else {
				log.Println("✅ Alert email sent")
			}
		}

		time.Sleep(10 * time.Second)
	}
}
