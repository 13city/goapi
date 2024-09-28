package main

import (
    "github.com/gin-gonic/gin"
    // Import your component packages
    "github.com/13city/goapi/portscanner"
    "github.com/13city/goapi/proxy"
    "github.com/13city/goapi/encryption"
    "github.com/13city/goapi/loganalysis"
    "github.com/13city/goapi/filetransfer"
    "github.com/13city/goapi/passwordstrength"
)

func main() {
    router := gin.Default()

    // Port Scanner Endpoint
    router.GET("/portscanner", portscanner.Handler)

    // Proxy Endpoint
    router.GET("/proxy", proxy.Handler)

    // Encryption Endpoints
    router.GET("/encrypt", encryption.EncryptHandler)
    router.GET("/decrypt", encryption.DecryptHandler)

    // Log Analysis Endpoint
    router.POST("/loganalysis", loganalysis.Handler)

    // File Transfer Endpoints
    router.POST("/upload", filetransfer.UploadHandler)
    router.GET("/download", filetransfer.DownloadHandler)

    // Password Strength Endpoint
    router.GET("/passwordstrength", passwordstrength.Handler)

    // Start the server
    router.Run(":8080")
}
