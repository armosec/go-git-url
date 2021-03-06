package giturl

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewGitURL(t *testing.T) {
	{ // parse github
		const githubURL = "https://github.com/armosec/go-git-url"
		gh, err := NewGitURL(githubURL)
		assert.NoError(t, err)
		assert.Equal(t, "github", gh.GetProvider())
		assert.Equal(t, "armosec", gh.GetOwnerName())
		assert.Equal(t, "go-git-url", gh.GetRepoName())
		assert.Equal(t, "", gh.GetBranchName())
		assert.Equal(t, githubURL, gh.GetURL().String())
	}
	{ // parse github
		const githubURL = "git@github.com:armosec/go-git-url.git"
		gh, err := NewGitURL(githubURL)
		assert.NoError(t, err)
		assert.Equal(t, "github", gh.GetProvider())
		assert.Equal(t, "armosec", gh.GetOwnerName())
		assert.Equal(t, "go-git-url", gh.GetRepoName())
		assert.Equal(t, "", gh.GetBranchName())
		assert.Equal(t, "https://github.com/armosec/go-git-url", gh.GetURL().String())
	}
	{ // parse azure
		const azureURL = "https://dev.azure.com/dwertent/ks-testing-public/_git/ks-testing-public"
		az, err := NewGitURL(azureURL)
		assert.NoError(t, err)
		assert.NoError(t, err)
		assert.Equal(t, "azure", az.GetProvider())
		assert.Equal(t, "dwertent", az.GetOwnerName())
		assert.Equal(t, "ks-testing-public", az.GetRepoName())
		assert.Equal(t, "", az.GetBranchName())
		assert.Equal(t, "", az.GetPath())
		assert.Equal(t, azureURL, az.GetURL().String())
	}
	{ // parse azure
		const azureURL = "git@ssh.dev.azure.com:v3/dwertent/ks-testing-public/ks-testing-public"
		az, err := NewGitURL(azureURL)
		assert.NoError(t, err)
		assert.NoError(t, err)
		assert.Equal(t, "azure", az.GetProvider())
		assert.Equal(t, "dwertent", az.GetOwnerName())
		assert.Equal(t, "ks-testing-public", az.GetRepoName())
		assert.Equal(t, "", az.GetBranchName())
		assert.Equal(t, "", az.GetPath())
		assert.Equal(t, "https://dev.azure.com/dwertent/ks-testing-public/_git/ks-testing-public", az.GetURL().String())
	}

}
func TestNewGitAPI(t *testing.T) {
	fileText := "https://raw.githubusercontent.com/armosec/go-git-url/master/files/file0.text"
	var gitURL IGitAPI
	var err error
	{
		gitURL, err = NewGitAPI("https://github.com/armosec/go-git-url")
		assert.NoError(t, err)

		files, err := gitURL.ListFilesNamesWithExtension([]string{"yaml", "json"})
		assert.NoError(t, err)
		assert.Equal(t, 7, len(files))
	}

	{
		gitURL, err = NewGitAPI("https://github.com/armosec/go-git-url")
		assert.NoError(t, err)

		files, errM := gitURL.DownloadFilesWithExtension([]string{"text"})
		assert.Equal(t, 0, len(errM))
		assert.Equal(t, 1, len(files))
		assert.Equal(t, "name=file0", string(files[fileText]))

	}

	{
		gitURL, err = NewGitAPI(fileText)
		assert.NoError(t, err)

		files, errM := gitURL.DownloadFilesWithExtension([]string{"text"})
		assert.Equal(t, 0, len(errM))
		assert.Equal(t, 1, len(files))
		assert.Equal(t, "name=file0", string(files[fileText]))
	}

	{
		gitURL, err = NewGitAPI(fileText)
		assert.NoError(t, err)

		files, errM := gitURL.DownloadAllFiles()
		assert.Equal(t, 0, len(errM))
		assert.Equal(t, 1, len(files))
		assert.Equal(t, "name=file0", string(files[fileText]))
	}

	{
		gitURL, err := NewGitAPI("https://github.com/armosec/go-git-url/tree/master/files")
		assert.NoError(t, err)

		files, errM := gitURL.DownloadFilesWithExtension([]string{"text"})
		assert.Equal(t, 0, len(errM))
		assert.Equal(t, 1, len(files))
		assert.Equal(t, "name=file0", string(files[fileText]))

	}

	{
		gitURL, err = NewGitAPI("https://github.com/armosec/go-git-url/blob/master/files/file0.text")
		assert.NoError(t, err)

		files, errM := gitURL.DownloadFilesWithExtension([]string{"text"})
		assert.Equal(t, 0, len(errM))
		assert.Equal(t, 1, len(files))
		assert.Equal(t, "name=file0", string(files[fileText]))

	}
}
