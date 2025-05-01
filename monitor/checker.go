package monitor

import (
    "fmt"
    "gopkg.in/gomail.v2"
    "github.com/shirou/gopsutil/v3/cpu"
    "github.com/shirou/gopsutil/v3/disk"
    "github.com/shirou/gopsutil/v3/mem"
    "time"
    "math/rand"
    "go-monitoring-system"
)

func CheckAndAlert(config main.Config) {
    cpuUsage, memUsage, diskUsage := getStats(config.Mock)

    if cpuUsage > config.CPUThreshold || memUsage > config.MemoryThreshold || diskUsage > config.DiskThreshold {
        body := fmt.Sprintf("⚠️ High Usage Detected:\nCPU: %.2f%%\nMemory: %.2f%%\nDisk: %.2f%%", cpuUsage, memUsage, diskUsage)
        sendEmail(body, config.EmailConfig)
    } else {
        fmt.Printf("✅ OK | CPU: %.2f%%, MEM: %.2f%%, DISK: %.2f%%\n", cpuUsage, memUsage, diskUsage)
    }
}

func getStats(mock bool) (float64, float64, float64) {
    if mock {
        rand.Seed(time.Now().UnixNano())
        return rand.Float64()*100, rand.Float64()*100, rand.Float64()*100
    }

    cpuPercents, _ := cpu.Percent(0, false)
    memStat, _ := mem.VirtualMemory()
    diskStat, _ := disk.Usage("/")

    return cpuPercents[0], memStat.UsedPercent, diskStat.UsedPercent
}

func sendEmail(body string, email main.EmailSettings) {
    m := gomail.NewMessage()
    m.SetHeader("From", email.From)
    m.SetHeader("To", email.To)
    m.SetHeader("Subject", "🚨 System Resource Alert")
    m.SetBody("text/plain", body)

    d := gomail.NewDialer(email.SMTPHost, email.SMTPPort, email.From, email.Password)

    if err := d.DialAndSend(m); err != nil {
        fmt.Println("❌ Failed to send email:", err)
    } else {
        fmt.Println("📧 Alert sent successfully!")
    }
}
