package config

func Init() error {
	err := ReadeConfig()
	if err != nil {
		return err
	}
	return nil
}
