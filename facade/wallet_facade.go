package facade

import "fmt"

type WalletFacade struct {
	account      *account
	wallet       *wallet
	securityCode *securityCode
	notification *notification
}

func NewWalletFacade(accountID string, code int) *WalletFacade {
	fmt.Println("开始创建账户")
	walletFacacde := &WalletFacade{
		account:      newAccount(accountID),
		securityCode: newSecurityCode(code),
		wallet:       newWallet(),
		notification: &notification{},
	}
	fmt.Println("账户创建成功")
	return walletFacacde
}

func (w *WalletFacade) AddMoneyToWallet(accountID string, securityCode int, amount int) error {
	fmt.Println("增加金额到钱包")
	err := w.account.checkAccount(accountID)
	if err != nil {
		return err
	}
	err = w.securityCode.checkCode(securityCode)
	if err != nil {
		return err
	}
	w.wallet.creditBalance(amount)
	w.notification.CreditNotification()
	return nil
}

func (w *WalletFacade) DeductMoneyFromWallet(accountID string, securityCode int, amount int) error {
	fmt.Println("从钱包中减少金额")
	err := w.account.checkAccount(accountID)
	if err != nil {
		return err
	}

	err = w.securityCode.checkCode(securityCode)
	if err != nil {
		return err
	}
	err = w.wallet.debitBalance(amount)
	if err != nil {
		return err
	}
	w.notification.DebitNotification()
	fmt.Println("从钱包中减少金额成功")
	return nil
}

type account struct {
	name string
}

func newAccount(accountName string) *account {
	return &account{
		name: accountName,
	}
}

func (a *account) checkAccount(accountName string) error {
	if a.name != accountName {
		return fmt.Errorf("用户名不正确")
	}
	fmt.Println("账户验证成功")
	return nil
}

type securityCode struct {
	code int
}

func newSecurityCode(code int) *securityCode {
	return &securityCode{
		code: code,
	}
}

func (s *securityCode) checkCode(incomingCode int) error {
	if s.code != incomingCode {
		return fmt.Errorf("安全码不正确")
	}
	fmt.Println("安全码正确")
	return nil
}

type wallet struct {
	balance int
}

func newWallet() *wallet {
	return &wallet{
		balance: 0,
	}
}

func (w *wallet) creditBalance(amount int) {
	w.balance += amount
	fmt.Println("增加金额成功")
}

func (w *wallet) debitBalance(amount int) error {
	if w.balance < amount {
		return fmt.Errorf("金额不够")
	}
	fmt.Println("减少金额成功")
	w.balance = w.balance - amount
	return nil
}

type notification struct {
}

func (n *notification) CreditNotification() {
	fmt.Println("你的账户金额增加了!")
}

func (n *notification) DebitNotification() {
	fmt.Println("你的账户金额减少了!")
}
