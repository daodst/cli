package core

type AccountClient struct {
	key *SecretKey
}

func (this *AccountClient) CreateAccountFromPriv(priv string) (*CosmosWallet, error) {
	return this.key.CreateAccountFromPriv(priv)
}

func (this *AccountClient) CreateAccountFromSeed(seed string) (acc *CosmosWallet, err error) {
	return this.key.CreateAccountFromSeed(seed)
}

func (this *AccountClient) CreateSeedWord() (mnemonic string, err error) {
	return this.key.CreateSeedWord()
}
