package wallet

import "testing"

func TestWallet(t *testing.T) {
	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))

		assertBalance(t, &wallet, Bitcoin(10))
	})

	t.Run("withdraw with funds", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(10))

		assertNoError(t, err)
		assertBalance(t, &wallet, Bitcoin(10))
	})

	t.Run("withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{balance: startingBalance}
		err := wallet.Withdraw(Bitcoin(100))

		assertError(t, err, ErrInsufficientFunds)
		assertBalance(t, &wallet, startingBalance)
	})
}

func assertBalance(t testing.TB, wallet *Wallet, expected Bitcoin) {
	t.Helper()

	got := wallet.Balance()
	if got != expected {
		t.Errorf("got %s expected %s", got, expected)
	}
}

func assertError(t testing.TB, got, expected error) {
	t.Helper()

	if got == nil {
		t.Fatal("expected an error but didn't get one")
	}

	if got != expected {
		t.Errorf("got %q, expected %q", got, expected)
	}
}

func assertNoError(t testing.TB, got error) {
	t.Helper()

	if got != nil {
		t.Fatalf("got error but didn't expect one. Got %q", got)
	}
}
