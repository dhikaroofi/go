package app

import "github.com/dhikaroofi/go/app/controller"

func (a *Mojoo) setRouters() {
	a.Post("/login", a.guest(controller.AuthLogin))
	a.Get("/merchant/list", a.guest(controller.GetListMerchants))
	a.Get("/merchant/{merchant_id}/outlet/list", a.guest(controller.GetOutletList))
	a.Get("/report/merchant/{merchant_id}/omzet", a.guest(controller.GetDailyReportMerchant))
	a.Get("/report/outlet/{outlet_id}/omzet", a.guest(controller.GetDailyReportOutlet))
}
