package common

var AccountIDs map[string]string = make(map[string]string)

func InitAccountIDs() error {
	accounts, err := Client.GetAccounts()
	if err != nil {
		return err
	}

	for _, account := range accounts {
		AccountIDs[account.Currency] = account.ID
	}

	return nil
}
