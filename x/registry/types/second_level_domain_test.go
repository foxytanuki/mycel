package types

import (
	"testing"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

type DomainTest struct {
	SecondLevelDomain SecondLevelDomain
	DomainParent      TopLevelDomain
	DomainPrice       sdk.Coins
}

type WalletRecordTest struct {
	WalletRecordType    string
	Address             string
	IsInvalidRecordType bool
	IsInvalidValue      bool
}

func TestDomainValidate(t *testing.T) {
	testCases := []struct {
		domain          SecondLevelDomain
		expDomainParent TopLevelDomain
		expErr          error
	}{
		// Valid domains
		{
			domain:          SecondLevelDomain{Name: "foo", Parent: "myc"},
			expDomainParent: TopLevelDomain{Name: "myc"},
		},
		// Invalid name
		{domain: SecondLevelDomain{Name: ".foo", Parent: "myc"},
			expErr: errorsmod.Wrapf(ErrInvalidDomainName, ".foo"),
		},
		{domain: SecondLevelDomain{Name: "", Parent: "myc"},
			expErr: errorsmod.Wrapf(ErrInvalidDomainName, ""),
		},
		{domain: SecondLevelDomain{Name: "bar.foo", Parent: "myc"},
			expErr: errorsmod.Wrapf(ErrInvalidDomainName, "bar.foo"),
		},
		{domain: SecondLevelDomain{Name: ".", Parent: "myc"},
			expErr: errorsmod.Wrapf(ErrInvalidDomainName, "."),
		},
		{domain: SecondLevelDomain{Name: "##", Parent: "myc"},
			expErr: errorsmod.Wrapf(ErrInvalidDomainName, "##"),
		},
		// Invalid parent
		{
			domain: SecondLevelDomain{Name: "foo", Parent: ".##"},
			expErr: errorsmod.Wrapf(ErrInvalidDomainParent, ".##"),
		},
		{
			domain: SecondLevelDomain{Name: "foo", Parent: ".myc"},
			expErr: errorsmod.Wrapf(ErrInvalidDomainParent, ".myc"),
		},
		{
			domain: SecondLevelDomain{Name: "foo", Parent: ".foo.myc"},
			expErr: errorsmod.Wrapf(ErrInvalidDomainParent, ".foo.myc"),
		},
	}

	for _, tc := range testCases {
		err := tc.domain.Validate()
		if tc.expErr == nil {
			require.Nil(t, err)

			// Check domain parent
			parent := tc.domain.ParseParent()
			require.Equal(t, tc.expDomainParent.Name, parent)
			// TODO: review test case
			// require.Equal(t, tc.expDomainParent.Name, name)
			// require.Equal(t, tc.expDomainParent.Parent, parent)

		} else {
			require.EqualError(t, err, tc.expErr.Error())
		}

	}

}

func TestDomainUpdateWalletRecord(t *testing.T) {
	testCases := []struct {
		walletRecordType string
		address          string
		expErr           error
	}{
		// Valid wallet records
		{walletRecordType: "BITCOIN_MAINNET_MAINNET", address: "1BvBMSEYstWetqTFn5Au4m4GFg7xJaNVN2"},
		{walletRecordType: "BITCOIN_MAINNET_MAINNET", address: "3J98t1WpEZ73CNmQviecrnyiWrnqRhWNLy"},
		{walletRecordType: "BITCOIN_MAINNET_MAINNET", address: "bc1qar0srrr7xfkvy5l643lydnw9re59gtzzwf5mdq"},

		// EVM
		{walletRecordType: "ETHEREUM_MAINNET_MAINNET", address: "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266"},
		{walletRecordType: "ETHEREUM_TESTNET_GOERLI", address: "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266"},
		{walletRecordType: "POLYGON_MAINNET_MAINNET", address: "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266"},
		{walletRecordType: "POLYGON_TESTNET_MUMBAI", address: "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266"},
		{walletRecordType: "BNB_MAINNET_MAINNET", address: "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266"},
		{walletRecordType: "BNB_TESTNET_TESTNET", address: "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266"},
		{walletRecordType: "AVALANCHE_MAINNET_CCHAIN", address: "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266"},
		{walletRecordType: "AVALANCHE_TESTNET_FUJI", address: "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266"},
		{walletRecordType: "GNOSIS_MAINNET_MAINNET", address: "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266"},
		{walletRecordType: "GNOSIS_TESTNET_CHIADO", address: "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266"},
		{walletRecordType: "OPTIMISM_MAINNET_MAINNET", address: "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266"},
		{walletRecordType: "OPTIMISM_TESTNET_GOERLI", address: "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266"},
		{walletRecordType: "ARBITRUM_MAINNET_MAINNET", address: "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266"},
		{walletRecordType: "ARBITRUM_TESTNET_GOERLI", address: "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266"},
		{walletRecordType: "SHARDEUM_BETANET_SPHINX", address: "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266"},
		{walletRecordType: "ZETA_TESTNET_ATHENS", address: "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266"},

		// Others
		{walletRecordType: "APTOS_MAINNET_MAINNET", address: "0xeeff357ea5c1a4e7bc11b2b17ff2dc2dcca69750bfef1e1ebcaccf8c8018175b"},
		{walletRecordType: "SOLANA_MAINNET_MAINNET", address: "HN7cABqLq46Es1jh92dQQisAq662SmxELLLsHHe4YWrH"},

		// Invalid record type
		{
			walletRecordType: "ETHEREUM_TESTNET_MUMBAI", address: "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266",
			expErr: errorsmod.Wrapf(ErrInvalidWalletRecordType, "ETHEREUM_TESTNET_MUMBAI"),
		},
		{
			walletRecordType: "ETHEREUM", address: "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266",
			expErr: errorsmod.Wrapf(ErrInvalidWalletRecordType, "ETHEREUM"),
		},
		// Invalid address
		{
			walletRecordType: "ETHEREUM_TESTNET_GOERLI", address: "0xf9Fd6e51aad88F6F4ce6aB8827279cffFb92266",
			expErr: errorsmod.Wrapf(ErrInvalidWalletAddress, "ETHEREUM 0xf9Fd6e51aad88F6F4ce6aB8827279cffFb92266"),
		},
		{
			walletRecordType: "ETHEREUM_TESTNET_GOERLI", address: "cosmos1jyc4rrtz5f93n80uuj378dq7x3v7z09j0h6dqx",
			expErr: errorsmod.Wrapf(ErrInvalidWalletAddress, "ETHEREUM cosmos1jyc4rrtz5f93n80uuj378dq7x3v7z09j0h6dqx"),
		},
		{
			walletRecordType: "SOLANA_MAINNET_MAINNET", address: "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266",
			expErr: errorsmod.Wrapf(ErrInvalidWalletAddress, "SOLANA 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266"),
		},
		{
			walletRecordType: "BITCOIN_MAINNET_MAINNET", address: "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266",
			expErr: errorsmod.Wrapf(ErrInvalidWalletAddress, "BITCOIN 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266"),
		},
	}
	for _, tc := range testCases {
		domain := SecondLevelDomain{Name: "foo", Parent: "myc"}
		err := domain.UpdateWalletRecord(tc.walletRecordType, tc.address)
		if tc.expErr == nil {
			require.Nil(t, err)
			require.Equal(t, tc.address, domain.Records[tc.walletRecordType].GetWalletRecord().GetValue())
		} else {
			require.EqualError(t, err, tc.expErr.Error())
		}
	}
}

func TestDomainUpdateDnsRecord(t *testing.T) {
	testCases := []struct {
		dnsRecordType string
		value         string
		expErr        error
	}{
		// Valid wallet records
		{dnsRecordType: "A", value: "10.0.0.1"},
		{dnsRecordType: "A", value: "192.168.0.1"},
		{dnsRecordType: "AAAA", value: "2001:0db8:85a3:0000:0000:8a2e:0370:7334"},
		{dnsRecordType: "CNAME", value: "example.com."},

		// Invalid record type
		{
			dnsRecordType: "FOO", value: "192.168.0.1",
			expErr: errorsmod.Wrapf(ErrInvalidDnsRecordType, "FOO"),
		},
		{
			dnsRecordType: "BAR", value: "192.168.0.1",
			expErr: errorsmod.Wrapf(ErrInvalidDnsRecordType, "BAR"),
		},
		// Invalid value
		{
			dnsRecordType: "A", value: "2001:0db8:85a3:0000:0000:8a2e:0370:7334",
			expErr: errorsmod.Wrapf(ErrInvalidDnsRecordValue, "IPV4 2001:0db8:85a3:0000:0000:8a2e:0370:7334"),
		},
		{
			dnsRecordType: "AAAA", value: "192.168.0.1",
			expErr: errorsmod.Wrapf(ErrInvalidDnsRecordValue, "IPV6 192.168.0.1"),
		},
	}
	for _, tc := range testCases {
		domain := SecondLevelDomain{Name: "foo", Parent: "myc"}
		err := domain.UpdateDnsRecord(tc.dnsRecordType, tc.value)
		if tc.expErr == nil {
			require.Nil(t, err)
			require.Equal(t, tc.value, domain.Records[tc.dnsRecordType].GetDnsRecord().GetValue())
		} else {
			require.EqualError(t, err, tc.expErr.Error())
		}
	}
}
