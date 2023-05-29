package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Book struct {
	Id          int
	Title       string
	Author      string
	ReleaseYear string
	Pages       int
}

var books []Book

var fileName string = "data.csv"

func main() {
	loadDataFromCSV(fileName)

	for {
		fmt.Println("==== Book Data Management ====")
		fmt.Println("1. View All Books")
		fmt.Println("2. Add New Book")
		fmt.Println("3. Update Book")
		fmt.Println("4. Delete Book")
		fmt.Println("5. Exit")

		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			viewAllBooks()
		case 2:
			addNewBook()
		case 3:
			updateBook()
		case 4:
			deleteBook()
		case 5:
			os.Exit(0)
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}

// func createFile(fileName string) {
// 	f, err := os.Create(fileName)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	} else {
// 		fmt.Println("File", fileName, "berhasil dibuat.")
// 	}
// 	defer f.Close()
// }

func loadDataFromCSV(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening CSV file:", err)
		return
	}
	defer file.Close()

	books = nil // reset book slice

	f, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		record := strings.Split(scanner.Text(), ",")
		id, _ := strconv.Atoi(record[0])
		pages, _ := strconv.Atoi(record[4])

		// Check if book with ID already exists
		bookIndex := findBookIndexByID(id)
		if bookIndex != -1 {
			fmt.Println("Buku dengan Id:", id, "tidak ada!")
		} else {
			// Book doesn't exist, create new book
			book := Book{
				Id:          id,
				Title:       record[1],
				Author:      record[2],
				ReleaseYear: record[3],
				Pages:       pages,
			}
			books = append(books, book)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading CSV record:", err)
		return
	}
}

func saveDataToCSV(fileName string) {

	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	for _, book := range books {
		row := strconv.Itoa(book.Id) + "," + book.Title + "," + book.Author + "," + book.ReleaseYear + "," + strconv.Itoa(book.Pages) + "\n"
		file.WriteString(row)
	}
	fmt.Println("Data berhasil disimpan ke file data.csv")
}

func viewAllBooks() {
	if len(books) == 0 {
		fmt.Println("No books available.")
		return
	}

	for _, book := range books {
		fmt.Printf("Book Id: %d\n", book.Id)
		fmt.Printf("Book Title: %s\n", book.Title)
		fmt.Printf("Book Author: %s\n", book.Author)
		fmt.Printf("Book ReleaseYear: %s\n", book.ReleaseYear)
		fmt.Printf("Book Pages: %d\n", book.Pages)
		fmt.Println()
	}
}

func addNewBook() {
	var newBook Book

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter Book Details:")
	fmt.Print("Book Id: ")
	scanner.Scan()
	newBook.Id, _ = strconv.Atoi(scanner.Text())

	fmt.Print("Book Title: ")
	scanner.Scan()
	newBook.Title = scanner.Text()

	fmt.Print("Book Author: ")
	scanner.Scan()
	newBook.Author = scanner.Text()

	fmt.Print("Book ReleaseYear: ")
	scanner.Scan()
	newBook.ReleaseYear = scanner.Text()

	fmt.Print("Book Pages: ")
	scanner.Scan()
	pages, _ := strconv.Atoi(scanner.Text())
	newBook.Pages = pages

	books = append(books, newBook)
	saveDataToCSV(fileName)

	fmt.Println("Book added successfully.")
}

func updateBook() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter Book Id to update: ")
	var bookId int
	scanner.Scan()
	bookId, _ = strconv.Atoi(scanner.Text())

	bookIndex := findBookIndexByID(bookId)
	if bookIndex == -1 {
		fmt.Println("Book not found.")
		return
	}

	var updatedBook Book

	fmt.Println("Enter Updated Book Details:")
	fmt.Print("Book Title: ")
	scanner.Scan()
	updatedBook.Title = scanner.Text()

	fmt.Print("Book Author: ")
	scanner.Scan()
	updatedBook.Author = scanner.Text()

	fmt.Print("Book ReleaseYear: ")
	scanner.Scan()
	updatedBook.ReleaseYear = scanner.Text()

	fmt.Print("Book Pages: ")
	scanner.Scan()
	pages, _ := strconv.Atoi(scanner.Text())
	updatedBook.Pages = pages

	books[bookIndex].Title = updatedBook.Title
	books[bookIndex].Author = updatedBook.Author
	books[bookIndex].ReleaseYear = updatedBook.ReleaseYear
	books[bookIndex].Pages = updatedBook.Pages
	saveDataToCSV(fileName)

	fmt.Println("Book updated successfully.")
}

func deleteBook() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter Book ID to delete: ")
	var bookId int
	scanner.Scan()
	bookId, _ = strconv.Atoi(scanner.Text())

	bookkIndex := findBookIndexByID(bookId)
	if bookkIndex == -1 {
		fmt.Println("Book not found.")
		return
	} else {
		books = append(books[:bookkIndex], books[bookkIndex+1:]...)
	}
	saveDataToCSV(fileName)

	fmt.Println("Book deleted successfully.")
}

func findBookIndexByID(id int) int {
	for i, book := range books {
		if book.Id == id {
			return i
		}
	}
	return -1
}
