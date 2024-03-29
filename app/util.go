package main

import (
	"encoding/binary"
	"fmt"
	"strconv"
	"strings"
)

func int32ToBytes(i int32) []byte {
	res := make([]byte, 4)
	binary.BigEndian.PutUint32(res, uint32(i))

	return res
}

func ParseIP(ip string) []byte {
	var response []byte
	elements := strings.Split(ip, ".")

	for _, element := range elements {
		num, err := strconv.Atoi(element)
		if err != nil {
			fmt.Println("Error converting string to int: ", err)
			return []byte{}
		}
		response = append(response, byte(num))
	}
	return response
}

func TypeNameToValue(s string) byte {
	switch s {
	case "A":
		return 1
	case "NS":
		return 2
	case "MD":
		return 3
	case "MF":
		return 4
	case "CNAME":
		return 5
	case "SOA":
		return 6
	case "MB":
		return 7
	case "MG":
		return 8
	case "MR":
		return 9
	case "NULL":
		return 10
	case "WKS":
		return 11
	case "PTR":
		return 12
	case "HINFO":
		return 13
	case "MINFO":
		return 14
	case "MX":
		return 15
	case "TXT":
		return 16
	default:
		return 0
	}
}

func ClassNameToValue(s string) byte {
	switch s {
	case "IN":
		return 1
	case "CS":
		return 2
	case "CH":
		return 3
	case "HS":
		return 4
	default:
		return 0
	}
}

func labelSequence(domain string) []byte {
	labels := strings.Split(domain, ".")
	var sequence []byte
	for _, label := range labels {
		sequence = append(sequence, byte(len(label)))
		sequence = append(sequence, label...)
	}
	sequence = append(sequence, '\x00')
	return sequence
}

// this function return the domain name and the offset of the next byte
func parseDomainName(data []byte, offset int) (string, int) {
	var res strings.Builder

	for {
		length := int(data[offset])
		offset++

		if length >= 192 {
			// pointerOffset := ((int(data[offset-1]) & 0x3F) << 8) + int(data[offset])
			pointerOffset := int(binary.BigEndian.Uint16(data[offset:offset+2])&0x3f) + 12

			subdomain, _ := parseDomainName(data, pointerOffset)

			res.WriteString(subdomain)
			offset++
			break
		}

		if offset >= len(data) || length == 0 {
			break
		}

		if res.Len() > 0 && length > 1 {
			res.WriteByte('.')
		}

		label := data[offset : offset+length]
		res.Write(label)
		offset += length
	}

	return res.String(), offset
}

func intToBytes(n int) []byte {
	val := make([]byte, 2)

	binary.BigEndian.PutUint16(val, uint16(n))

	return val
}

func toBytes(s string) []byte {
	return []byte(s)
}

func GetIPLength(ip string) int {
	processed_ip := strings.ReplaceAll(ip, ".", "")

	return len(processed_ip)
}
