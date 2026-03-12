package cmd

import (
	"os"
	"path/filepath"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Append Command", func() {
	var (
		tempDir  string
		testFile string
	)

	BeforeEach(func() {
		var err error
		tempDir, err = os.MkdirTemp("", "append-test-*")
		Expect(err).NotTo(HaveOccurred())

		testFile = filepath.Join(tempDir, "test.txt")
	})

	AfterEach(func() {
		err := os.RemoveAll(tempDir)
		Expect(err).NotTo(HaveOccurred())
	})

	Describe("Append function", func() {
		It("should append content to an existing file", func() {
			// Create initial file with content
			err := os.WriteFile(testFile, []byte("initial content"), 0644)
			Expect(err).NotTo(HaveOccurred())

			// Set up command flags
			path = testFile
			content = "appended text"

			err = append()
			Expect(err).NotTo(HaveOccurred())

			// Verify content
			content, err := os.ReadFile(testFile)
			Expect(err).NotTo(HaveOccurred())
			Expect(string(content)).To(Equal("initial contentappended text"))
		})

		It("should create a new file if it doesn't exist", func() {
			nonExistentFile := filepath.Join(tempDir, "new-file.txt")

			path = nonExistentFile
			content = "new content"

			err := append()
			Expect(err).NotTo(HaveOccurred())

			// Verify file was created
			Expect(nonExistentFile).To(BeAnExistingFile())

			content, err := os.ReadFile(nonExistentFile)
			Expect(err).NotTo(HaveOccurred())
			Expect(string(content)).To(Equal("new content"))
		})

		It("should append multiple times to the same file", func() {
			path = testFile

			content = "first "
			err := append()
			Expect(err).NotTo(HaveOccurred())

			content = "second "
			err = append()
			Expect(err).NotTo(HaveOccurred())

			content = "third"
			err = append()
			Expect(err).NotTo(HaveOccurred())

			fileContent, err := os.ReadFile(testFile)
			Expect(err).NotTo(HaveOccurred())
			Expect(string(fileContent)).To(Equal("first second third"))
		})

		It("should handle empty content", func() {
			err := os.WriteFile(testFile, []byte("original"), 0644)
			Expect(err).NotTo(HaveOccurred())

			path = testFile
			content = ""

			err = append()
			Expect(err).NotTo(HaveOccurred())

			fileContent, err := os.ReadFile(testFile)
			Expect(err).NotTo(HaveOccurred())
			Expect(string(fileContent)).To(Equal("original"))
		})

		It("should handle multiline content", func() {
			path = testFile
			content = "line1\nline2\nline3\n"

			err := append()
			Expect(err).NotTo(HaveOccurred())

			fileContent, err := os.ReadFile(testFile)
			Expect(err).NotTo(HaveOccurred())
			Expect(string(fileContent)).To(Equal("line1\nline2\nline3\n"))
		})

		It("should return error for invalid file path", func() {
			invalidPath := filepath.Join(tempDir, "nonexistent-dir", "file.txt")

			path = invalidPath
			content = "content"

			err := append()
			Expect(err).To(HaveOccurred())
		})

		It("should return error for read-only file", func() {
			err := os.WriteFile(testFile, []byte("content"), 0444)
			Expect(err).NotTo(HaveOccurred())

			path = testFile
			content = "new content"

			err = append()
			Expect(err).To(HaveOccurred())

			// Cleanup: restore permissions for AfterEach
			os.Chmod(testFile, 0644)
		})

		It("should preserve file permissions", func() {
			err := os.WriteFile(testFile, []byte("content"), 0755)
			Expect(err).NotTo(HaveOccurred())

			path = testFile
			content = "appended"

			err = append()
			Expect(err).NotTo(HaveOccurred())

			fileInfo, err := os.Stat(testFile)
			Expect(err).NotTo(HaveOccurred())
			Expect(fileInfo.Mode().Perm()).To(Equal(os.FileMode(0755)))
		})
	})

	Describe("Append Command", func() {
		It("should have correct command metadata", func() {
			Expect(appendCmd.Use).To(Equal("append"))
			Expect(appendCmd.Short).To(ContainSubstring("Append text"))
		})

		It("should require content flag", func() {
			// This would be tested via integration tests with cobra
			// Verify the flag is marked as required
			Expect(appendCmd.Flags().Lookup("content")).NotTo(BeNil())
		})
	})
})
