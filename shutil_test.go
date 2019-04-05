package shutil

import (
	"bytes"
	"io/ioutil"
	"os"
	"testing"
)

func filesMatch(src, dst string) (bool, error) {
	srcContents, err := ioutil.ReadFile(src)
	if err != nil {
		return false, err
	}

	dstContents, err := ioutil.ReadFile(dst)
	if err != nil {
		return false, err
	}

	if bytes.Compare(srcContents, dstContents) != 0 {
		return false, nil
	}
	return true, nil
}

func TestSameFileError(t *testing.T) {
	_, err := Copy("testdata/testfile", "testdata/testfile", false)
	_, ok := err.(*SameFileError)
	if !ok {
		t.Error(err)
	}
}

func TestCopyFile(t *testing.T) {
	// clear out existing files if they exist
	os.Remove("testdata/testfile3")
	defer os.Remove("testdata/testfile3")

	err := CopyFile("testdata/testfile", "testdata/testfile3", false)
	if err != nil {
		t.Error(err)
		return
	}

	match, err := filesMatch("testdata/testfile", "testdata/testfile3")
	if err != nil {
		t.Error(err)
		return
	}
	if !match {
		t.Fail()
		return
	}

	// And again without clearing the files
	err = CopyFile("testdata/testfile2", "testdata/testfile3", false)
	if err != nil {
		t.Error(err)
		return
	}

	match2, err := filesMatch("testdata/testfile2", "testdata/testfile3")
	if err != nil {
		t.Error(err)
		return
	}

	if !match2 {
		t.Fail()
		return
	}
}

func TestCopy(t *testing.T) {
	// clear out existing files if they exist
	os.Remove("testdata/testfile3")
	defer os.Remove("testdata/testfile3")

	_, err := Copy("testdata/testfile", "testdata/testfile3", false)
	if err != nil {
		t.Error(err)
		return
	}

	match, err := filesMatch("testdata/testfile", "testdata/testfile3")
	if err != nil {
		t.Error(err)
		return
	}
	if !match {
		t.Fail()
		return
	}

	// And again without clearing the files
	_, err = Copy("testdata/testfile2", "testdata/testfile3", false)
	if err != nil {
		t.Error(err)
		return
	}

	match2, err := filesMatch("testdata/testfile2", "testdata/testfile3")
	if err != nil {
		t.Error(err)
		return
	}

	if !match2 {
		t.Fail()
		return
	}
}

func TestCopyTree(t *testing.T) {
	// clear out existing files if they exist
	os.RemoveAll("testdata/testdir3")
	defer os.RemoveAll("testdata/testdir3")

	err := CopyTree("testdata/testdir", "testdata/testdir3", nil)
	if err != nil {
		t.Error(err)
		return
	}

	match, err := filesMatch("testdata/testdir/file1", "testdata/testdir3/file1")
	if err != nil {
		t.Error(err)
		return
	}
	if !match {
		t.Fail()
		return
	}

	// And again without clearing the files
	err = CopyTree("testdata/testdir", "testdata/testdir3", nil)
	_, ok := err.(*AlreadyExistsError)
	if err == nil || !ok {
		t.Errorf("Expected an AlreadyExistsError but got: %s", err)
		return
	}

	// And again without clearing the files (but AllowExistingDestination set)
	err = CopyTree("testdata/testdir", "testdata/testdir3", &CopyTreeOptions{AllowExistingDestination: true})
	if err != nil {
		t.Error(err)
		return
	}

	match2, err := filesMatch("testdata/testdir/file1", "testdata/testdir/file1")
	if err != nil {
		t.Error(err)
		return
	}

	if !match2 {
		t.Fail()
		return
	}
}
