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

	err := loadDataFromCSV(fileName)
	if err != nil {
		panic(err)
	}

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
			err := viewAllBooks()
			if err != nil {
				fmt.Println(err)
			}
		case 2:
			err := addNewBook()
			if err != nil {
				fmt.Println(err)
			}
		case 3:
			err := updateBook()
			if err != nil {
				fmt.Println(err)
			}
		case 4:
			err := deleteBook()
			if err != nil {
				fmt.Println(err)
			}
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

func loadDataFromCSV(fileName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return fmt.Errorf("error opening csv file: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		record := strings.Split(scanner.Text(), ",")
		id, _ := strconv.Atoi(record[0])
		pages, _ := strconv.Atoi(record[4])

		book := Book{
			Id:          id,
			Title:       record[1],
			Author:      record[2],
			ReleaseYear: record[3],
			Pages:       pages,
		}
		books = append(books, book)
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error opening csv file: %w", err)
	}
	return nil
}

func saveDataToCSV(fileName string) error {
	file, err := os.Create(fileName)
	if err != nil {
		return fmt.Errorf("error opening csv file: %w", err)
	}
	defer file.Close()

	for _, book := range books {
		row := strconv.Itoa(book.Id) + "," + book.Title + "," + book.Author + "," + book.ReleaseYear + "," + strconv.Itoa(book.Pages) + "\n"
		file.WriteString(row)
	}
	fmt.Println("Data successfully saved to data.csv")
	return nil
}

func viewAllBooks() error {
	if len(books) == 0 {
		return fmt.Errorf("no books available")
	}

	for _, book := range books {
		fmt.Printf("Book Id: %d\n", book.Id)
		fmt.Printf("Book Title: %s\n", book.Title)
		fmt.Printf("Book Author: %s\n", book.Author)
		fmt.Printf("Book ReleaseYear: %s\n", book.ReleaseYear)
		fmt.Printf("Book Pages: %d\n", book.Pages)
		fmt.Println()
	}

	return nil
}

func addNewBook() error {
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

	fmt.Print("Are you sure want to add this book (y/n)?")
	scanner.Scan()
	choice := scanner.Text()
	if strings.ToLower(choice) == "y" {
		// dari pada kalian buat seperti ini, lebih baik kalian buat idnya auto increment
		// dengan cara, tidak usah kalian beri inputan untuk id dan idnya kalian cek dari id data sebelumnya di tambahkan 1
		_, err := findBookIndexByID(newBook.Id)
		if err != nil {
			books = append(books, newBook)
		} else {
			return fmt.Errorf("Book with id: %d already exist", newBook.Id)
		}

		fmt.Println("Book added successfully.")
		err = saveDataToCSV(fileName)
		if err != nil {
			return err
		}

	} else if strings.ToLower(choice) == "n" {
		fmt.Println("Data is not saved")
	} else {
		fmt.Println("Invalid choice. Please try again.")
		addNewBook()
	}

	return nil
}

func updateBook() error {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter Book Id to update: ")
	var bookId int
	scanner.Scan()
	bookId, _ = strconv.Atoi(scanner.Text())

	bookIndex, err := findBookIndexByID(bookId)
	if err != nil {
		return err
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

	fmt.Print("Are you sure want to update this book (y/n)?")
	scanner.Scan()
	choice := scanner.Text()
	if strings.ToLower(choice) == "y" {
		if updatedBook.Title != "" {
			books[bookIndex].Title = updatedBook.Title
		}
		if updatedBook.Author != "" {
			books[bookIndex].Author = updatedBook.Author
		}
		if updatedBook.ReleaseYear != "" {
			books[bookIndex].ReleaseYear = updatedBook.ReleaseYear
		}
		if updatedBook.Pages != 0 {
			books[bookIndex].Pages = updatedBook.Pages
		}

		fmt.Println("Book updated successfully.")
		err = saveDataToCSV(fileName)
		if err != nil {
			return err
		}
	} else if strings.ToLower(choice) == "n" {
		fmt.Println("Data is not updated")
	} else {
		fmt.Println("Invalid choice. Please try again.")
		updateBook()
	}

	return nil
}

func deleteBook() error {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter Book ID to delete: ")
	var bookId int
	scanner.Scan()
	bookId, _ = strconv.Atoi(scanner.Text())

	fmt.Print("Are you sure want to delete this book (y/n)?")
	scanner.Scan()
	choice := scanner.Text()
	if strings.ToLower(choice) == "y" {
		bookIndex, err := findBookIndexByID(bookId)
		if err != nil {
			return err
		} else {
			books = append(books[:bookIndex], books[bookIndex+1:]...)
		}
		fmt.Println("Book deleted successfully.")
		err = saveDataToCSV(fileName)
		if err != nil {
			return err
		}
	} else if strings.ToLower(choice) == "n" {
		fmt.Println("Data is not deleted")
	} else {
		fmt.Println("Invalid choice. Please try again.")
		deleteBook()
	}

	return nil
}

func findBookIndexByID(id int) (int, error) {
	for i, book := range books {
		if book.Id == id {
			return i, nil
		}
	}
	return 0, fmt.Errorf("id: %d not found", id)
}
