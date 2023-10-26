package usecase

// func GetAllPayments(payment *models.Payment) (res []payloads.GetAllPaymentsResponse, err error) {
// 	payments, err := repositories.GetAllPayments(payment)
// 	if err != nil {
// 		return res, err
// 	}

// 	res = []payloads.GetAllPaymentsResponse{}
// 	for i, payment := range payments {
// 		res = append(res, payloads.GetAllPaymentsResponse{
// 			ID:  payment.ID,
// 			Spp: payment.Spp.ID,
// 		})
// 	}
// }
