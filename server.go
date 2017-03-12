package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"net"
	"strconv"
	"time"

	"github.com/sandp125/ticketing/goticketing"
)

func check(err error, message string) {
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", message)
}

func random(min, max int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return rand.Intn(max-min) + min
}

func main() {
	theatre := goticketing.GetTheatreMoq()
	cafeteria := goticketing.GetCafeteriaMoq()

	windowID := 0
	receiptCh := make(chan goticketing.Receipt)
	personCh := make(chan int)

	ln, err := net.Listen("tcp", ":8080")
	check(err, "Server is ready.")

	for {
		conn, err := ln.Accept()
		check(err, "Accepted connection.")
		windowID++

		go func() {
			buf := bufio.NewReader(conn)

			switch windowID {
			case 1:
				go func() {
					conn.Write([]byte("Welcome - Window 1\r\n"))
					for {
						_, err := buf.ReadString('\n')

						if err != nil {
							fmt.Printf("Client disconnected.\n")
							break
						}

						conn.Write([]byte("<<<Purchased Ticket - Window 1>>>\r\n"))

						var screenNumber, showNumber, personNumber int
						screenNumber = random(0, 5)
						showNumber = random(0, 4)
						personNumber = random(1, 1000)

						if theatre.Screens[screenNumber].Shows[showNumber].NoOfSeats > 0 {
							theatre.Screens[screenNumber].Shows[showNumber].NoOfSeats--
							conn.Write([]byte("Screen Number: " + strconv.Itoa(screenNumber+1) + "			Show Number: " + strconv.Itoa(showNumber+1) + "			Person Number: " + strconv.Itoa(personNumber) + ".\r\n"))
							receipt := goticketing.Receipt{
								ScreenNumber: screenNumber + 1,
								ShowNumber:   showNumber + 1,
								PersonNumber: personNumber,
							}
							receiptCh <- receipt
							personCh <- personNumber
							cafeteria.TotalWaterSold++
							cafeteria.TotalPopcornSold++
							theatre.TotalTicketsSold++
						} else {
							conn.Write([]byte("Screen Number: " + strconv.Itoa(screenNumber+1) + "			Show Number: " + strconv.Itoa(showNumber+1) + " is at full capacity.\r\n"))
						}
					}
				}()
			case 2:
				go func() {
					conn.Write([]byte("Welcome - Window 2\r\n"))
					for {
						_, err := buf.ReadString('\n')

						if err != nil {
							fmt.Printf("Client disconnected.\n")
							break
						}

						conn.Write([]byte("<<<Purchased Ticket - Window 2>>>\r\n"))

						var screenNumber, showNumber, personNumber int
						screenNumber = random(0, 5)
						showNumber = random(0, 4)
						personNumber = random(1, 1000)

						if theatre.Screens[screenNumber].Shows[showNumber].NoOfSeats > 0 {
							theatre.Screens[screenNumber].Shows[showNumber].NoOfSeats--
							conn.Write([]byte("Screen Number: " + strconv.Itoa(screenNumber+1) + "			Show Number: " + strconv.Itoa(showNumber+1) + "			Person Number: " + strconv.Itoa(personNumber) + ".\r\n"))
							receipt := goticketing.Receipt{
								ScreenNumber: screenNumber + 1,
								ShowNumber:   showNumber + 1,
								PersonNumber: personNumber,
							}
							receiptCh <- receipt
							theatre.TotalTicketsSold++
						} else {
							conn.Write([]byte("Screen Number: " + strconv.Itoa(screenNumber+1) + "			Show Number: " + strconv.Itoa(showNumber+1) + " is at full capacity.\r\n"))
						}
					}
				}()
			case 3:
				go func() {
					conn.Write([]byte("Welcome - Cafeteria\r\n"))
					for l := range personCh {
						if l%2 == 0 {
							if cafeteria.NoOfSodas > 0 {
								conn.Write([]byte("Person: " + strconv.Itoa(l) + " exchanged Water with Soda.\r\n"))
								cafeteria.NoOfSodas--
								cafeteria.TotalSodasSold++
								cafeteria.TotalWaterSold--
								conn.Write([]byte("Soda left: " + strconv.Itoa(cafeteria.NoOfSodas) + ".\r\n"))
							} else {
								conn.Write([]byte("No Soda left in the Cafeteria.\r\n"))
							}
						}
					}
				}()
				go func() {
					_, err := buf.ReadString('\n')

					if err != nil {
						fmt.Printf("Client disconnected.\n")
					}
				}()
			case 4:
				conn.Write([]byte("Welcome - Sales\r\n"))
				ticker := time.NewTicker(time.Millisecond * 30000)
				go func() {
					for l := range receiptCh {
						conn.Write([]byte("Ticket Purchased by Person: " + strconv.Itoa(l.PersonNumber) + ".\r\n"))
					}
				}()
				go func() {
					for t := range ticker.C {
						conn.Write([]byte("Sales Report at " + t.String() + ".\r\n"))
						conn.Write([]byte("----------------------------------------------------------------------\r\n"))
						conn.Write([]byte("Tickets Sold: " + strconv.Itoa(theatre.TotalTicketsSold) + ".\r\n"))
						conn.Write([]byte("Soda Sold: " + strconv.Itoa(cafeteria.TotalSodasSold) + ".\r\n"))
						conn.Write([]byte("Water Sold: " + strconv.Itoa(cafeteria.TotalWaterSold) + ".\r\n"))
						conn.Write([]byte("Popcorn Sold: " + strconv.Itoa(cafeteria.TotalPopcornSold) + ".\r\n"))
					}
				}()
				go func() {
					_, err := buf.ReadString('\n')

					if err != nil {
						fmt.Printf("Client disconnected.\n")
					}
				}()
			}
		}()
	}
}
