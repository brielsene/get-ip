package controllers

import (
	"get-ip/dtos"
	"net"
	"os"

	"github.com/gin-gonic/gin"
)

// func GetIp(c *gin.Context) {
// 	// Obtém o hostname local
// 	hostname, err := net.LookupHost("localhost")
// 	if err != nil {
// 		c.JSON(400, gin.H{"error": err.Error()})
// 		return
// 	}

// 	// Obtém o primeiro IP associado ao hostname
// 	ipAddresses, err := net.LookupIP("localhost")
// 	if err != nil {
// 		c.JSON(400, gin.H{"error": err.Error()})
// 		return
// 	}

// 	// Verifica se existem IPs e pega o primeiro
// 	if len(ipAddresses) == 0 {
// 		c.JSON(400, gin.H{"error": "nenhum endereço IP encontrado"})
// 		return
// 	}
// 	ip := ipAddresses[0].String()

// 	// Concatena o hostname e o IP
// 	hostip := hostname[0] + " - " + ip

// 	// Retorna o resultado em JSON
// 	c.JSON(200, gin.H{"hostname-ip": hostip})
// }

func GetIp(c *gin.Context) {
	hostname, err := os.Hostname()
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	addrs, err := net.LookupIP(hostname)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// Filtrar para obter o primeiro endereço IPv4
	var ipv4 string
	for _, addr := range addrs {
		if addr.To4() != nil {
			ipv4 = addr.String()
			break
		}
	}

	if ipv4 == "" {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	dto := dtos.DtoRede{}
	dto.Hostanme = hostname
	dto.Ip = ipv4
	c.JSON(200, dto)
}
