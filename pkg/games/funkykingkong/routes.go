
// pkg/games/funkykingkong/routes.go
package funkykingkong

import (
	"fmt"
	"strings"

	"github.com/JILI-GAMES/b_backend_games11/pkg/common/rng"
	"github.com/JILI-GAMES/b_backend_games11/pkg/common/settings"
	"github.com/gofiber/fiber/v2"
)

// RouteGroup holds the dependencies for game routes
type RouteGroup struct {
	RNGProd      *rng.Client
	SettingsProd *settings.Client
	RNGTest      *rng.Client
	SettingsTest *settings.Client
}

// NewRouteGroup creates a new route group for funky king kong game
func NewRouteGroup(rngProd *rng.Client, settingsProd *settings.Client, rngTest *rng.Client, settingsTest *settings.Client) *RouteGroup {
	return &RouteGroup{
		RNGProd:      rngProd,
		SettingsProd: settingsProd,
		RNGTest:      rngTest,
		SettingsTest: settingsTest,
	}
}

// Helper to select the correct clients per request
func (rg *RouteGroup) getClientsForRequest(c *fiber.Ctx) (*rng.Client, *settings.Client) {
	origin := c.Get("Origin")
	fmt.Printf("Origin: %s\n", origin)
	if len(origin) > 0 && (strings.Contains(strings.ToLower(origin), "test")) {
		return rg.RNGTest, rg.SettingsTest
	}
	return rg.RNGProd, rg.SettingsProd
}

// Register registers the funky king kong game routes
func (rg *RouteGroup) Register(app *fiber.App) {
	app.Post("/spin/funkykingkong", rg.SpinHandler)
}
