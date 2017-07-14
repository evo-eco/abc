package logout

import (
	"fmt"
	"github.com/appbaseio/abc/appbase/session"
	"os"
)

// UserLogout log outs a user
func UserLogout() error {
	err := session.DeleteUserSession()
	if err != nil {
		return err
	}
	// remove env var
	os.Unsetenv("ABC_TOKEN")
	fmt.Println("Logged out successfully")
	return nil
}