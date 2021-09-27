package echo

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

func Bind(c echo.Context, postData interface{}) error {
	if err := c.Bind(postData); err != nil {
		return fmt.Errorf("failed to bind the data: %v", err)
	}

	if err := c.Validate(postData); err != nil {
		return fmt.Errorf("failed to verify the data: %v", err)
	}

	return nil
}
