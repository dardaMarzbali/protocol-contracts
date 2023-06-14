// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package zetatokenconsumertrident

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// ZetaTokenConsumerTridentMetaData contains all meta data concerning the ZetaTokenConsumerTrident contract.
var ZetaTokenConsumerTridentMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"zetaToken_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"uniswapV3Router_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"WETH9Address_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"poolFactory_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ErrorSendingETH\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InputCantBeZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyError\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"name\":\"EthExchangedForZeta\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"name\":\"TokenExchangedForZeta\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"name\":\"ZetaExchangedForEth\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"name\":\"ZetaExchangedForToken\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minAmountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"zetaTokenAmount\",\"type\":\"uint256\"}],\"name\":\"getEthFromZeta\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minAmountOut\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"outputToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"zetaTokenAmount\",\"type\":\"uint256\"}],\"name\":\"getTokenFromZeta\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minAmountOut\",\"type\":\"uint256\"}],\"name\":\"getZetaFromEth\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minAmountOut\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"inputToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"inputTokenAmount\",\"type\":\"uint256\"}],\"name\":\"getZetaFromToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poolFactory\",\"outputs\":[{\"internalType\":\"contractConcentratedLiquidityPoolFactory\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tridentRouter\",\"outputs\":[{\"internalType\":\"contractIPoolRouter\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"zetaToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x6101006040523480156200001257600080fd5b5060405162002af538038062002af5833981810160405281019062000038919062000245565b600073ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff161480620000a05750600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16145b80620000d85750600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16145b80620001105750600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16145b1562000148576040517fe6c4247b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8373ffffffffffffffffffffffffffffffffffffffff1660a08173ffffffffffffffffffffffffffffffffffffffff1660601b815250508273ffffffffffffffffffffffffffffffffffffffff1660c08173ffffffffffffffffffffffffffffffffffffffff1660601b815250508073ffffffffffffffffffffffffffffffffffffffff1660e08173ffffffffffffffffffffffffffffffffffffffff1660601b815250508173ffffffffffffffffffffffffffffffffffffffff1660808173ffffffffffffffffffffffffffffffffffffffff1660601b81525050505050506200030a565b6000815190506200023f81620002f0565b92915050565b60008060008060808587031215620002625762000261620002eb565b5b600062000272878288016200022e565b945050602062000285878288016200022e565b935050604062000298878288016200022e565b9250506060620002ab878288016200022e565b91505092959194509250565b6000620002c482620002cb565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600080fd5b620002fb81620002b7565b81146200030757600080fd5b50565b60805160601c60a05160601c60c05160601c60e05160601c6126e162000414600039600081816102e0015281816106ca015281816107d601528181610b0c01528181610cde015281816111a8015261129401526000818161042c0152818161064f01528181610a1201528181610c2301528181610e4901528181610f3e015281816110ed01526114f00152600081816102b401528181610521015281816107a601528181610bd901528181610c4501528181610c9101528181610da3015281816110a30152818161110f0152818161115b015261147b0152600081816102930152818161069e0152818161078501528181610cb20152818161117c015261126301526126e16000f3fe6080604052600436106100745760003560e01c80634219dc401161004e5780634219dc401461011857806354c49a2a1461014357806364b5528a14610180578063a53fb10b146101ab5761007b565b8063013b2ff81461008057806321e093b1146100b05780632405620a146100db5761007b565b3661007b57005b600080fd5b61009a60048036038101906100959190611bd5565b6101e8565b6040516100a791906122b5565b60405180910390f35b3480156100bc57600080fd5b506100c561051f565b6040516100d29190612080565b60405180910390f35b3480156100e757600080fd5b5061010260048036038101906100fd9190611c15565b610543565b60405161010f91906122b5565b60405180910390f35b34801561012457600080fd5b5061012d610b0a565b60405161013a91906121a0565b60405180910390f35b34801561014f57600080fd5b5061016a60048036038101906101659190611c7c565b610b2e565b60405161017791906122b5565b60405180910390f35b34801561018c57600080fd5b50610195610f3c565b6040516101a291906121bb565b60405180910390f35b3480156101b757600080fd5b506101d260048036038101906101cd9190611c15565b610f60565b6040516101df91906122b5565b60405180910390f35b60008073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161415610250576040517fe6c4247b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600034141561028b576040517fb813f54900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000806102d87f00000000000000000000000000000000000000000000000000000000000000007f0000000000000000000000000000000000000000000000000000000000000000611602565b9150915060007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166371a258128484600060016040518563ffffffff1660e01b815260040161033f94939291906120c4565b60006040518083038186803b15801561035757600080fd5b505afa15801561036b573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f820116820180604052508101906103949190611ccf565b905060006040518060c00160405280600073ffffffffffffffffffffffffffffffffffffffff168152602001348152602001878152602001836000815181106103e0576103df6124cd565b5b602002602001015173ffffffffffffffffffffffffffffffffffffffff1681526020018873ffffffffffffffffffffffffffffffffffffffff16815260200160001515815250905060007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663c07f5c3234846040518363ffffffff1660e01b8152600401610484919061229a565b6020604051808303818588803b15801561049d57600080fd5b505af11580156104b1573d6000803e3d6000fd5b50505050506040513d601f19601f820116820180604052508101906104d69190611d45565b90507f87890b0a30955b1db16cc894fbe24779ae05d9f337bfe8b6dfc0809b5bf9da1134826040516105099291906122d0565b60405180910390a1809550505050505092915050565b7f000000000000000000000000000000000000000000000000000000000000000081565b60008073ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff1614806105ab5750600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16145b156105e2576040517fe6c4247b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600082141561061d576040517fb813f54900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61064a3330848673ffffffffffffffffffffffffffffffffffffffff16611652909392919063ffffffff16565b6106957f0000000000000000000000000000000000000000000000000000000000000000838573ffffffffffffffffffffffffffffffffffffffff166116db9092919063ffffffff16565b6000806106c2857f0000000000000000000000000000000000000000000000000000000000000000611602565b9150915060007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166371a258128484600060016040518563ffffffff1660e01b815260040161072994939291906120c4565b60006040518083038186803b15801561074157600080fd5b505afa158015610755573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f8201168201806040525081019061077e9190611ccf565b90506107ca7f00000000000000000000000000000000000000000000000000000000000000007f0000000000000000000000000000000000000000000000000000000000000000611602565b809350819450505060007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166371a258128585600060016040518563ffffffff1660e01b815260040161083594939291906120c4565b60006040518083038186803b15801561084d57600080fd5b505afa158015610861573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f8201168201806040525081019061088a9190611ccf565b90506000600267ffffffffffffffff8111156108a9576108a86124fc565b5b6040519080825280602002602001820160405280156108d75781602001602082028036833780820191505090505b509050826000815181106108ee576108ed6124cd565b5b60200260200101518160008151811061090a576109096124cd565b5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff168152505081600081518110610958576109576124cd565b5b602002602001015181600181518110610974576109736124cd565b5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff168152505060006040518060c001604052808a73ffffffffffffffffffffffffffffffffffffffff1681526020018981526020018b81526020018381526020018c73ffffffffffffffffffffffffffffffffffffffff16815260200160001515815250905060007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663363a9dba836040518263ffffffff1660e01b8152600401610a699190612278565b602060405180830381600087803b158015610a8357600080fd5b505af1158015610a97573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610abb9190611d45565b90507f017190d3d99ee6d8dd0604ef1e71ff9802810c6485f57c9b2ed6169848dd119f8a8a83604051610af093929190612169565b60405180910390a180975050505050505050949350505050565b7f000000000000000000000000000000000000000000000000000000000000000081565b60008073ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff161415610b96576040517fe6c4247b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000821415610bd1576040517fb813f54900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610c1e3330847f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16611652909392919063ffffffff16565b610c897f0000000000000000000000000000000000000000000000000000000000000000837f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166116db9092919063ffffffff16565b600080610cd67f00000000000000000000000000000000000000000000000000000000000000007f0000000000000000000000000000000000000000000000000000000000000000611602565b9150915060007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166371a258128484600060016040518563ffffffff1660e01b8152600401610d3d94939291906120c4565b60006040518083038186803b158015610d5557600080fd5b505afa158015610d69573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f82011682018060405250810190610d929190611ccf565b905060006040518060c001604052807f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16815260200187815260200188815260200183600081518110610dfd57610dfc6124cd565b5b602002602001015173ffffffffffffffffffffffffffffffffffffffff1681526020018973ffffffffffffffffffffffffffffffffffffffff16815260200160011515815250905060007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663c07f5c32836040518263ffffffff1660e01b8152600401610ea0919061229a565b602060405180830381600087803b158015610eba57600080fd5b505af1158015610ece573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ef29190611d45565b90507f74e171117e91660f493740924d8bad0caf48dc4fbccb767fb05935397a2c17ae8782604051610f259291906122d0565b60405180910390a180955050505050509392505050565b7f000000000000000000000000000000000000000000000000000000000000000081565b60008060009054906101000a900460ff1615610fa8576040517f29f745a700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60016000806101000a81548160ff021916908315150217905550600073ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff1614806110295750600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16145b15611060576040517fe6c4247b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600082141561109b576040517fb813f54900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6110e83330847f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16611652909392919063ffffffff16565b6111537f0000000000000000000000000000000000000000000000000000000000000000837f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166116db9092919063ffffffff16565b6000806111a07f00000000000000000000000000000000000000000000000000000000000000007f0000000000000000000000000000000000000000000000000000000000000000611602565b9150915060007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166371a258128484600060016040518563ffffffff1660e01b815260040161120794939291906120c4565b60006040518083038186803b15801561121f57600080fd5b505afa158015611233573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f8201168201806040525081019061125c9190611ccf565b90506112887f000000000000000000000000000000000000000000000000000000000000000087611602565b809350819450505060007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166371a258128585600060016040518563ffffffff1660e01b81526004016112f394939291906120c4565b60006040518083038186803b15801561130b57600080fd5b505afa15801561131f573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f820116820180604052508101906113489190611ccf565b90506000600267ffffffffffffffff811115611367576113666124fc565b5b6040519080825280602002602001820160405280156113955781602001602082028036833780820191505090505b509050826000815181106113ac576113ab6124cd565b5b6020026020010151816000815181106113c8576113c76124cd565b5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff168152505081600081518110611416576114156124cd565b5b602002602001015181600181518110611432576114316124cd565b5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff168152505060006040518060c001604052807f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1681526020018981526020018b81526020018381526020018c73ffffffffffffffffffffffffffffffffffffffff16815260200160001515815250905060007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663363a9dba836040518263ffffffff1660e01b81526004016115479190612278565b602060405180830381600087803b15801561156157600080fd5b505af1158015611575573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906115999190611d45565b90507f0a7cb8f6e1d29e616c1209a3f418c17b3a9137005377f6dd072217b1ede2573b8a8a836040516115ce93929190612169565b60405180910390a18097505050505050505060008060006101000a81548160ff021916908315150217905550949350505050565b6000808273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff1610156116445783839150915061164b565b8284915091505b9250929050565b6116d5846323b872dd60e01b85858560405160240161167393929190612109565b604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050611839565b50505050565b6000811480611774575060008373ffffffffffffffffffffffffffffffffffffffff1663dd62ed3e30856040518363ffffffff1660e01b815260040161172292919061209b565b60206040518083038186803b15801561173a57600080fd5b505afa15801561174e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906117729190611d45565b145b6117b3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016117aa90612258565b60405180910390fd5b6118348363095ea7b360e01b84846040516024016117d2929190612140565b604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050611839565b505050565b600061189b826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65648152508573ffffffffffffffffffffffffffffffffffffffff166119009092919063ffffffff16565b90506000815111156118fb57808060200190518101906118bb9190611d18565b6118fa576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016118f190612238565b60405180910390fd5b5b505050565b606061190f8484600085611918565b90509392505050565b60608247101561195d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611954906121f8565b60405180910390fd5b6000808673ffffffffffffffffffffffffffffffffffffffff1685876040516119869190612069565b60006040518083038185875af1925050503d80600081146119c3576040519150601f19603f3d011682016040523d82523d6000602084013e6119c8565b606091505b50915091506119d9878383876119e5565b92505050949350505050565b60608315611a4857600083511415611a4057611a0085611a5b565b611a3f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611a3690612218565b60405180910390fd5b5b829050611a53565b611a528383611a7e565b5b949350505050565b6000808273ffffffffffffffffffffffffffffffffffffffff163b119050919050565b600082511115611a915781518083602001fd5b806040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611ac591906121d6565b60405180910390fd5b6000611ae1611adc8461231e565b6122f9565b90508083825260208201905082856020860282011115611b0457611b03612530565b5b60005b85811015611b345781611b1a8882611b53565b845260208401935060208301925050600181019050611b07565b5050509392505050565b600081359050611b4d81612666565b92915050565b600081519050611b6281612666565b92915050565b600082601f830112611b7d57611b7c61252b565b5b8151611b8d848260208601611ace565b91505092915050565b600081519050611ba58161267d565b92915050565b600081359050611bba81612694565b92915050565b600081519050611bcf81612694565b92915050565b60008060408385031215611bec57611beb61253a565b5b6000611bfa85828601611b3e565b9250506020611c0b85828601611bab565b9150509250929050565b60008060008060808587031215611c2f57611c2e61253a565b5b6000611c3d87828801611b3e565b9450506020611c4e87828801611bab565b9350506040611c5f87828801611b3e565b9250506060611c7087828801611bab565b91505092959194509250565b600080600060608486031215611c9557611c9461253a565b5b6000611ca386828701611b3e565b9350506020611cb486828701611bab565b9250506040611cc586828701611bab565b9150509250925092565b600060208284031215611ce557611ce461253a565b5b600082015167ffffffffffffffff811115611d0357611d02612535565b5b611d0f84828501611b68565b91505092915050565b600060208284031215611d2e57611d2d61253a565b5b6000611d3c84828501611b96565b91505092915050565b600060208284031215611d5b57611d5a61253a565b5b6000611d6984828501611bc0565b91505092915050565b6000611d7e8383611d8a565b60208301905092915050565b611d93816123b5565b82525050565b611da2816123b5565b82525050565b6000611db38261235a565b611dbd8185612388565b9350611dc88361234a565b8060005b83811015611df9578151611de08882611d72565b9750611deb8361237b565b925050600181019050611dcc565b5085935050505092915050565b611e0f816123c7565b82525050565b6000611e2082612365565b611e2a8185612399565b9350611e3a818560208601612469565b80840191505092915050565b611e4f816123fd565b82525050565b611e5e8161240f565b82525050565b611e6d81612421565b82525050565b611e7c81612433565b82525050565b6000611e8d82612370565b611e9781856123a4565b9350611ea7818560208601612469565b611eb08161253f565b840191505092915050565b6000611ec86026836123a4565b9150611ed382612550565b604082019050919050565b6000611eeb601d836123a4565b9150611ef68261259f565b602082019050919050565b6000611f0e602a836123a4565b9150611f19826125c8565b604082019050919050565b6000611f316036836123a4565b9150611f3c82612617565b604082019050919050565b600060c083016000830151611f5f6000860182611d8a565b506020830151611f72602086018261204b565b506040830151611f85604086018261204b565b5060608301518482036060860152611f9d8282611da8565b9150506080830151611fb26080860182611d8a565b5060a0830151611fc560a0860182611e06565b508091505092915050565b60c082016000820151611fe66000850182611d8a565b506020820151611ff9602085018261204b565b50604082015161200c604085018261204b565b50606082015161201f6060850182611d8a565b5060808201516120326080850182611d8a565b5060a082015161204560a0850182611e06565b50505050565b612054816123f3565b82525050565b612063816123f3565b82525050565b60006120758284611e15565b915081905092915050565b60006020820190506120956000830184611d99565b92915050565b60006040820190506120b06000830185611d99565b6120bd6020830184611d99565b9392505050565b60006080820190506120d96000830187611d99565b6120e66020830186611d99565b6120f36040830185611e64565b6121006060830184611e73565b95945050505050565b600060608201905061211e6000830186611d99565b61212b6020830185611d99565b612138604083018461205a565b949350505050565b60006040820190506121556000830185611d99565b612162602083018461205a565b9392505050565b600060608201905061217e6000830186611d99565b61218b602083018561205a565b612198604083018461205a565b949350505050565b60006020820190506121b56000830184611e46565b92915050565b60006020820190506121d06000830184611e55565b92915050565b600060208201905081810360008301526121f08184611e82565b905092915050565b6000602082019050818103600083015261221181611ebb565b9050919050565b6000602082019050818103600083015261223181611ede565b9050919050565b6000602082019050818103600083015261225181611f01565b9050919050565b6000602082019050818103600083015261227181611f24565b9050919050565b600060208201905081810360008301526122928184611f47565b905092915050565b600060c0820190506122af6000830184611fd0565b92915050565b60006020820190506122ca600083018461205a565b92915050565b60006040820190506122e5600083018561205a565b6122f2602083018461205a565b9392505050565b6000612303612314565b905061230f828261249c565b919050565b6000604051905090565b600067ffffffffffffffff821115612339576123386124fc565b5b602082029050602081019050919050565b6000819050602082019050919050565b600081519050919050565b600081519050919050565b600081519050919050565b6000602082019050919050565b600082825260208201905092915050565b600081905092915050565b600082825260208201905092915050565b60006123c0826123d3565b9050919050565b60008115159050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600061240882612445565b9050919050565b600061241a82612445565b9050919050565b600061242c826123f3565b9050919050565b600061243e826123f3565b9050919050565b600061245082612457565b9050919050565b6000612462826123d3565b9050919050565b60005b8381101561248757808201518184015260208101905061246c565b83811115612496576000848401525b50505050565b6124a58261253f565b810181811067ffffffffffffffff821117156124c4576124c36124fc565b5b80604052505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f416464726573733a20696e73756666696369656e742062616c616e636520666f60008201527f722063616c6c0000000000000000000000000000000000000000000000000000602082015250565b7f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000600082015250565b7f5361666545524332303a204552433230206f7065726174696f6e20646964206e60008201527f6f74207375636365656400000000000000000000000000000000000000000000602082015250565b7f5361666545524332303a20617070726f76652066726f6d206e6f6e2d7a65726f60008201527f20746f206e6f6e2d7a65726f20616c6c6f77616e636500000000000000000000602082015250565b61266f816123b5565b811461267a57600080fd5b50565b612686816123c7565b811461269157600080fd5b50565b61269d816123f3565b81146126a857600080fd5b5056fea26469706673582212204468ddc2d97337c8f733ef254edb6c213627014c4fa5f7dea76ca62adde499e064736f6c63430008070033",
}

// ZetaTokenConsumerTridentABI is the input ABI used to generate the binding from.
// Deprecated: Use ZetaTokenConsumerTridentMetaData.ABI instead.
var ZetaTokenConsumerTridentABI = ZetaTokenConsumerTridentMetaData.ABI

// ZetaTokenConsumerTridentBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ZetaTokenConsumerTridentMetaData.Bin instead.
var ZetaTokenConsumerTridentBin = ZetaTokenConsumerTridentMetaData.Bin

// DeployZetaTokenConsumerTrident deploys a new Ethereum contract, binding an instance of ZetaTokenConsumerTrident to it.
func DeployZetaTokenConsumerTrident(auth *bind.TransactOpts, backend bind.ContractBackend, zetaToken_ common.Address, uniswapV3Router_ common.Address, WETH9Address_ common.Address, poolFactory_ common.Address) (common.Address, *types.Transaction, *ZetaTokenConsumerTrident, error) {
	parsed, err := ZetaTokenConsumerTridentMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ZetaTokenConsumerTridentBin), backend, zetaToken_, uniswapV3Router_, WETH9Address_, poolFactory_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ZetaTokenConsumerTrident{ZetaTokenConsumerTridentCaller: ZetaTokenConsumerTridentCaller{contract: contract}, ZetaTokenConsumerTridentTransactor: ZetaTokenConsumerTridentTransactor{contract: contract}, ZetaTokenConsumerTridentFilterer: ZetaTokenConsumerTridentFilterer{contract: contract}}, nil
}

// ZetaTokenConsumerTrident is an auto generated Go binding around an Ethereum contract.
type ZetaTokenConsumerTrident struct {
	ZetaTokenConsumerTridentCaller     // Read-only binding to the contract
	ZetaTokenConsumerTridentTransactor // Write-only binding to the contract
	ZetaTokenConsumerTridentFilterer   // Log filterer for contract events
}

// ZetaTokenConsumerTridentCaller is an auto generated read-only Go binding around an Ethereum contract.
type ZetaTokenConsumerTridentCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZetaTokenConsumerTridentTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ZetaTokenConsumerTridentTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZetaTokenConsumerTridentFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ZetaTokenConsumerTridentFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZetaTokenConsumerTridentSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ZetaTokenConsumerTridentSession struct {
	Contract     *ZetaTokenConsumerTrident // Generic contract binding to set the session for
	CallOpts     bind.CallOpts             // Call options to use throughout this session
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// ZetaTokenConsumerTridentCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ZetaTokenConsumerTridentCallerSession struct {
	Contract *ZetaTokenConsumerTridentCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                   // Call options to use throughout this session
}

// ZetaTokenConsumerTridentTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ZetaTokenConsumerTridentTransactorSession struct {
	Contract     *ZetaTokenConsumerTridentTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                   // Transaction auth options to use throughout this session
}

// ZetaTokenConsumerTridentRaw is an auto generated low-level Go binding around an Ethereum contract.
type ZetaTokenConsumerTridentRaw struct {
	Contract *ZetaTokenConsumerTrident // Generic contract binding to access the raw methods on
}

// ZetaTokenConsumerTridentCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ZetaTokenConsumerTridentCallerRaw struct {
	Contract *ZetaTokenConsumerTridentCaller // Generic read-only contract binding to access the raw methods on
}

// ZetaTokenConsumerTridentTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ZetaTokenConsumerTridentTransactorRaw struct {
	Contract *ZetaTokenConsumerTridentTransactor // Generic write-only contract binding to access the raw methods on
}

// NewZetaTokenConsumerTrident creates a new instance of ZetaTokenConsumerTrident, bound to a specific deployed contract.
func NewZetaTokenConsumerTrident(address common.Address, backend bind.ContractBackend) (*ZetaTokenConsumerTrident, error) {
	contract, err := bindZetaTokenConsumerTrident(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ZetaTokenConsumerTrident{ZetaTokenConsumerTridentCaller: ZetaTokenConsumerTridentCaller{contract: contract}, ZetaTokenConsumerTridentTransactor: ZetaTokenConsumerTridentTransactor{contract: contract}, ZetaTokenConsumerTridentFilterer: ZetaTokenConsumerTridentFilterer{contract: contract}}, nil
}

// NewZetaTokenConsumerTridentCaller creates a new read-only instance of ZetaTokenConsumerTrident, bound to a specific deployed contract.
func NewZetaTokenConsumerTridentCaller(address common.Address, caller bind.ContractCaller) (*ZetaTokenConsumerTridentCaller, error) {
	contract, err := bindZetaTokenConsumerTrident(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ZetaTokenConsumerTridentCaller{contract: contract}, nil
}

// NewZetaTokenConsumerTridentTransactor creates a new write-only instance of ZetaTokenConsumerTrident, bound to a specific deployed contract.
func NewZetaTokenConsumerTridentTransactor(address common.Address, transactor bind.ContractTransactor) (*ZetaTokenConsumerTridentTransactor, error) {
	contract, err := bindZetaTokenConsumerTrident(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ZetaTokenConsumerTridentTransactor{contract: contract}, nil
}

// NewZetaTokenConsumerTridentFilterer creates a new log filterer instance of ZetaTokenConsumerTrident, bound to a specific deployed contract.
func NewZetaTokenConsumerTridentFilterer(address common.Address, filterer bind.ContractFilterer) (*ZetaTokenConsumerTridentFilterer, error) {
	contract, err := bindZetaTokenConsumerTrident(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ZetaTokenConsumerTridentFilterer{contract: contract}, nil
}

// bindZetaTokenConsumerTrident binds a generic wrapper to an already deployed contract.
func bindZetaTokenConsumerTrident(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ZetaTokenConsumerTridentMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZetaTokenConsumerTrident *ZetaTokenConsumerTridentRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZetaTokenConsumerTrident.Contract.ZetaTokenConsumerTridentCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZetaTokenConsumerTrident *ZetaTokenConsumerTridentRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZetaTokenConsumerTrident.Contract.ZetaTokenConsumerTridentTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZetaTokenConsumerTrident *ZetaTokenConsumerTridentRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZetaTokenConsumerTrident.Contract.ZetaTokenConsumerTridentTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZetaTokenConsumerTrident *ZetaTokenConsumerTridentCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZetaTokenConsumerTrident.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZetaTokenConsumerTrident *ZetaTokenConsumerTridentTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZetaTokenConsumerTrident.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZetaTokenConsumerTrident *ZetaTokenConsumerTridentTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZetaTokenConsumerTrident.Contract.contract.Transact(opts, method, params...)
}

// PoolFactory is a free data retrieval call binding the contract method 0x4219dc40.
//
// Solidity: function poolFactory() view returns(address)
func (_ZetaTokenConsumerTrident *ZetaTokenConsumerTridentCaller) PoolFactory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ZetaTokenConsumerTrident.contract.Call(opts, &out, "poolFactory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PoolFactory is a free data retrieval call binding the contract method 0x4219dc40.
//
// Solidity: function poolFactory() view returns(address)
func (_ZetaTokenConsumerTrident *ZetaTokenConsumerTridentSession) PoolFactory() (common.Address, error) {
	return _ZetaTokenConsumerTrident.Contract.PoolFactory(&_ZetaTokenConsumerTrident.CallOpts)
}

// PoolFactory is a free data retrieval call binding the contract method 0x4219dc40.
//
// Solidity: function poolFactory() view returns(address)
func (_ZetaTokenConsumerTrident *ZetaTokenConsumerTridentCallerSession) PoolFactory() (common.Address, error) {
	return _ZetaTokenConsumerTrident.Contract.PoolFactory(&_ZetaTokenConsumerTrident.CallOpts)
}

// TridentRouter is a free data retrieval call binding the contract method 0x64b5528a.
//
// Solidity: function tridentRouter() view returns(address)
func (_ZetaTokenConsumerTrident *ZetaTokenConsumerTridentCaller) TridentRouter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ZetaTokenConsumerTrident.contract.Call(opts, &out, "tridentRouter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TridentRouter is a free data retrieval call binding the contract method 0x64b5528a.
//
// Solidity: function tridentRouter() view returns(address)
func (_ZetaTokenConsumerTrident *ZetaTokenConsumerTridentSession) TridentRouter() (common.Address, error) {
	return _ZetaTokenConsumerTrident.Contract.TridentRouter(&_ZetaTokenConsumerTrident.CallOpts)
}

// TridentRouter is a free data retrieval call binding the contract method 0x64b5528a.
//
// Solidity: function tridentRouter() view returns(address)
func (_ZetaTokenConsumerTrident *ZetaTokenConsumerTridentCallerSession) TridentRouter() (common.Address, error) {
	return _ZetaTokenConsumerTrident.Contract.TridentRouter(&_ZetaTokenConsumerTrident.CallOpts)
}

// ZetaToken is a free data retrieval call binding the contract method 0x21e093b1.
//
// Solidity: function zetaToken() view returns(address)
func (_ZetaTokenConsumerTrident *ZetaTokenConsumerTridentCaller) ZetaToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ZetaTokenConsumerTrident.contract.Call(opts, &out, "zetaToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ZetaToken is a free data retrieval call binding the contract method 0x21e093b1.
//
// Solidity: function zetaToken() view returns(address)
func (_ZetaTokenConsumerTrident *ZetaTokenConsumerTridentSession) ZetaToken() (common.Address, error) {
	return _ZetaTokenConsumerTrident.Contract.ZetaToken(&_ZetaTokenConsumerTrident.CallOpts)
}

// ZetaToken is a free data retrieval call binding the contract method 0x21e093b1.
//
// Solidity: function zetaToken() view returns(address)
func (_ZetaTokenConsumerTrident *ZetaTokenConsumerTridentCallerSession) ZetaToken() (common.Address, error) {
	return _ZetaTokenConsumerTrident.Contract.ZetaToken(&_ZetaTokenConsumerTrident.CallOpts)
}

// GetEthFromZeta is a paid mutator transaction binding the contract method 0x54c49a2a.
//
// Solidity: function getEthFromZeta(address destinationAddress, uint256 minAmountOut, uint256 zetaTokenAmount) returns(uint256)
func (_ZetaTokenConsumerTrident *ZetaTokenConsumerTridentTransactor) GetEthFromZeta(opts *bind.TransactOpts, destinationAddress common.Address, minAmountOut *big.Int, zetaTokenAmount *big.Int) (*types.Transaction, error) {
	return _ZetaTokenConsumerTrident.contract.Transact(opts, "getEthFromZeta", destinationAddress, minAmountOut, zetaTokenAmount)
}

// GetEthFromZeta is a paid mutator transaction binding the contract method 0x54c49a2a.
//
// Solidity: function getEthFromZeta(address destinationAddress, uint256 minAmountOut, uint256 zetaTokenAmount) returns(uint256)
func (_ZetaTokenConsumerTrident *ZetaTokenConsumerTridentSession) GetEthFromZeta(destinationAddress common.Address, minAmountOut *big.Int, zetaTokenAmount *big.Int) (*types.Transaction, error) {
	return _ZetaTokenConsumerTrident.Contract.GetEthFromZeta(&_ZetaTokenConsumerTrident.TransactOpts, destinationAddress, minAmountOut, zetaTokenAmount)
}

// GetEthFromZeta is a paid mutator transaction binding the contract method 0x54c49a2a.
//
// Solidity: function getEthFromZeta(address destinationAddress, uint256 minAmountOut, uint256 zetaTokenAmount) returns(uint256)
func (_ZetaTokenConsumerTrident *ZetaTokenConsumerTridentTransactorSession) GetEthFromZeta(destinationAddress common.Address, minAmountOut *big.Int, zetaTokenAmount *big.Int) (*types.Transaction, error) {
	return _ZetaTokenConsumerTrident.Contract.GetEthFromZeta(&_ZetaTokenConsumerTrident.TransactOpts, destinationAddress, minAmountOut, zetaTokenAmount)
}

// GetTokenFromZeta is a paid mutator transaction binding the contract method 0xa53fb10b.
//
// Solidity: function getTokenFromZeta(address destinationAddress, uint256 minAmountOut, address outputToken, uint256 zetaTokenAmount) returns(uint256)
func (_ZetaTokenConsumerTrident *ZetaTokenConsumerTridentTransactor) GetTokenFromZeta(opts *bind.TransactOpts, destinationAddress common.Address, minAmountOut *big.Int, outputToken common.Address, zetaTokenAmount *big.Int) (*types.Transaction, error) {
	return _ZetaTokenConsumerTrident.contract.Transact(opts, "getTokenFromZeta", destinationAddress, minAmountOut, outputToken, zetaTokenAmount)
}

// GetTokenFromZeta is a paid mutator transaction binding the contract method 0xa53fb10b.
//
// Solidity: function getTokenFromZeta(address destinationAddress, uint256 minAmountOut, address outputToken, uint256 zetaTokenAmount) returns(uint256)
func (_ZetaTokenConsumerTrident *ZetaTokenConsumerTridentSession) GetTokenFromZeta(destinationAddress common.Address, minAmountOut *big.Int, outputToken common.Address, zetaTokenAmount *big.Int) (*types.Transaction, error) {
	return _ZetaTokenConsumerTrident.Contract.GetTokenFromZeta(&_ZetaTokenConsumerTrident.TransactOpts, destinationAddress, minAmountOut, outputToken, zetaTokenAmount)
}

// GetTokenFromZeta is a paid mutator transaction binding the contract method 0xa53fb10b.
//
// Solidity: function getTokenFromZeta(address destinationAddress, uint256 minAmountOut, address outputToken, uint256 zetaTokenAmount) returns(uint256)
func (_ZetaTokenConsumerTrident *ZetaTokenConsumerTridentTransactorSession) GetTokenFromZeta(destinationAddress common.Address, minAmountOut *big.Int, outputToken common.Address, zetaTokenAmount *big.Int) (*types.Transaction, error) {
	return _ZetaTokenConsumerTrident.Contract.GetTokenFromZeta(&_ZetaTokenConsumerTrident.TransactOpts, destinationAddress, minAmountOut, outputToken, zetaTokenAmount)
}

// GetZetaFromEth is a paid mutator transaction binding the contract method 0x013b2ff8.
//
// Solidity: function getZetaFromEth(address destinationAddress, uint256 minAmountOut) payable returns(uint256)
func (_ZetaTokenConsumerTrident *ZetaTokenConsumerTridentTransactor) GetZetaFromEth(opts *bind.TransactOpts, destinationAddress common.Address, minAmountOut *big.Int) (*types.Transaction, error) {
	return _ZetaTokenConsumerTrident.contract.Transact(opts, "getZetaFromEth", destinationAddress, minAmountOut)
}

// GetZetaFromEth is a paid mutator transaction binding the contract method 0x013b2ff8.
//
// Solidity: function getZetaFromEth(address destinationAddress, uint256 minAmountOut) payable returns(uint256)
func (_ZetaTokenConsumerTrident *ZetaTokenConsumerTridentSession) GetZetaFromEth(destinationAddress common.Address, minAmountOut *big.Int) (*types.Transaction, error) {
	return _ZetaTokenConsumerTrident.Contract.GetZetaFromEth(&_ZetaTokenConsumerTrident.TransactOpts, destinationAddress, minAmountOut)
}

// GetZetaFromEth is a paid mutator transaction binding the contract method 0x013b2ff8.
//
// Solidity: function getZetaFromEth(address destinationAddress, uint256 minAmountOut) payable returns(uint256)
func (_ZetaTokenConsumerTrident *ZetaTokenConsumerTridentTransactorSession) GetZetaFromEth(destinationAddress common.Address, minAmountOut *big.Int) (*types.Transaction, error) {
	return _ZetaTokenConsumerTrident.Contract.GetZetaFromEth(&_ZetaTokenConsumerTrident.TransactOpts, destinationAddress, minAmountOut)
}

// GetZetaFromToken is a paid mutator transaction binding the contract method 0x2405620a.
//
// Solidity: function getZetaFromToken(address destinationAddress, uint256 minAmountOut, address inputToken, uint256 inputTokenAmount) returns(uint256)
func (_ZetaTokenConsumerTrident *ZetaTokenConsumerTridentTransactor) GetZetaFromToken(opts *bind.TransactOpts, destinationAddress common.Address, minAmountOut *big.Int, inputToken common.Address, inputTokenAmount *big.Int) (*types.Transaction, error) {
	return _ZetaTokenConsumerTrident.contract.Transact(opts, "getZetaFromToken", destinationAddress, minAmountOut, inputToken, inputTokenAmount)
}

// GetZetaFromToken is a paid mutator transaction binding the contract method 0x2405620a.
//
// Solidity: function getZetaFromToken(address destinationAddress, uint256 minAmountOut, address inputToken, uint256 inputTokenAmount) returns(uint256)
func (_ZetaTokenConsumerTrident *ZetaTokenConsumerTridentSession) GetZetaFromToken(destinationAddress common.Address, minAmountOut *big.Int, inputToken common.Address, inputTokenAmount *big.Int) (*types.Transaction, error) {
	return _ZetaTokenConsumerTrident.Contract.GetZetaFromToken(&_ZetaTokenConsumerTrident.TransactOpts, destinationAddress, minAmountOut, inputToken, inputTokenAmount)
}

// GetZetaFromToken is a paid mutator transaction binding the contract method 0x2405620a.
//
// Solidity: function getZetaFromToken(address destinationAddress, uint256 minAmountOut, address inputToken, uint256 inputTokenAmount) returns(uint256)
func (_ZetaTokenConsumerTrident *ZetaTokenConsumerTridentTransactorSession) GetZetaFromToken(destinationAddress common.Address, minAmountOut *big.Int, inputToken common.Address, inputTokenAmount *big.Int) (*types.Transaction, error) {
	return _ZetaTokenConsumerTrident.Contract.GetZetaFromToken(&_ZetaTokenConsumerTrident.TransactOpts, destinationAddress, minAmountOut, inputToken, inputTokenAmount)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ZetaTokenConsumerTrident *ZetaTokenConsumerTridentTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZetaTokenConsumerTrident.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ZetaTokenConsumerTrident *ZetaTokenConsumerTridentSession) Receive() (*types.Transaction, error) {
	return _ZetaTokenConsumerTrident.Contract.Receive(&_ZetaTokenConsumerTrident.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ZetaTokenConsumerTrident *ZetaTokenConsumerTridentTransactorSession) Receive() (*types.Transaction, error) {
	return _ZetaTokenConsumerTrident.Contract.Receive(&_ZetaTokenConsumerTrident.TransactOpts)
}

// ZetaTokenConsumerTridentEthExchangedForZetaIterator is returned from FilterEthExchangedForZeta and is used to iterate over the raw logs and unpacked data for EthExchangedForZeta events raised by the ZetaTokenConsumerTrident contract.
type ZetaTokenConsumerTridentEthExchangedForZetaIterator struct {
	Event *ZetaTokenConsumerTridentEthExchangedForZeta // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZetaTokenConsumerTridentEthExchangedForZetaIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaTokenConsumerTridentEthExchangedForZeta)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ZetaTokenConsumerTridentEthExchangedForZeta)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ZetaTokenConsumerTridentEthExchangedForZetaIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaTokenConsumerTridentEthExchangedForZetaIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaTokenConsumerTridentEthExchangedForZeta represents a EthExchangedForZeta event raised by the ZetaTokenConsumerTrident contract.
type ZetaTokenConsumerTridentEthExchangedForZeta struct {
	AmountIn  *big.Int
	AmountOut *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterEthExchangedForZeta is a free log retrieval operation binding the contract event 0x87890b0a30955b1db16cc894fbe24779ae05d9f337bfe8b6dfc0809b5bf9da11.
//
// Solidity: event EthExchangedForZeta(uint256 amountIn, uint256 amountOut)
func (_ZetaTokenConsumerTrident *ZetaTokenConsumerTridentFilterer) FilterEthExchangedForZeta(opts *bind.FilterOpts) (*ZetaTokenConsumerTridentEthExchangedForZetaIterator, error) {

	logs, sub, err := _ZetaTokenConsumerTrident.contract.FilterLogs(opts, "EthExchangedForZeta")
	if err != nil {
		return nil, err
	}
	return &ZetaTokenConsumerTridentEthExchangedForZetaIterator{contract: _ZetaTokenConsumerTrident.contract, event: "EthExchangedForZeta", logs: logs, sub: sub}, nil
}

// WatchEthExchangedForZeta is a free log subscription operation binding the contract event 0x87890b0a30955b1db16cc894fbe24779ae05d9f337bfe8b6dfc0809b5bf9da11.
//
// Solidity: event EthExchangedForZeta(uint256 amountIn, uint256 amountOut)
func (_ZetaTokenConsumerTrident *ZetaTokenConsumerTridentFilterer) WatchEthExchangedForZeta(opts *bind.WatchOpts, sink chan<- *ZetaTokenConsumerTridentEthExchangedForZeta) (event.Subscription, error) {

	logs, sub, err := _ZetaTokenConsumerTrident.contract.WatchLogs(opts, "EthExchangedForZeta")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaTokenConsumerTridentEthExchangedForZeta)
				if err := _ZetaTokenConsumerTrident.contract.UnpackLog(event, "EthExchangedForZeta", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseEthExchangedForZeta is a log parse operation binding the contract event 0x87890b0a30955b1db16cc894fbe24779ae05d9f337bfe8b6dfc0809b5bf9da11.
//
// Solidity: event EthExchangedForZeta(uint256 amountIn, uint256 amountOut)
func (_ZetaTokenConsumerTrident *ZetaTokenConsumerTridentFilterer) ParseEthExchangedForZeta(log types.Log) (*ZetaTokenConsumerTridentEthExchangedForZeta, error) {
	event := new(ZetaTokenConsumerTridentEthExchangedForZeta)
	if err := _ZetaTokenConsumerTrident.contract.UnpackLog(event, "EthExchangedForZeta", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaTokenConsumerTridentTokenExchangedForZetaIterator is returned from FilterTokenExchangedForZeta and is used to iterate over the raw logs and unpacked data for TokenExchangedForZeta events raised by the ZetaTokenConsumerTrident contract.
type ZetaTokenConsumerTridentTokenExchangedForZetaIterator struct {
	Event *ZetaTokenConsumerTridentTokenExchangedForZeta // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZetaTokenConsumerTridentTokenExchangedForZetaIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaTokenConsumerTridentTokenExchangedForZeta)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ZetaTokenConsumerTridentTokenExchangedForZeta)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ZetaTokenConsumerTridentTokenExchangedForZetaIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaTokenConsumerTridentTokenExchangedForZetaIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaTokenConsumerTridentTokenExchangedForZeta represents a TokenExchangedForZeta event raised by the ZetaTokenConsumerTrident contract.
type ZetaTokenConsumerTridentTokenExchangedForZeta struct {
	Token     common.Address
	AmountIn  *big.Int
	AmountOut *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTokenExchangedForZeta is a free log retrieval operation binding the contract event 0x017190d3d99ee6d8dd0604ef1e71ff9802810c6485f57c9b2ed6169848dd119f.
//
// Solidity: event TokenExchangedForZeta(address token, uint256 amountIn, uint256 amountOut)
func (_ZetaTokenConsumerTrident *ZetaTokenConsumerTridentFilterer) FilterTokenExchangedForZeta(opts *bind.FilterOpts) (*ZetaTokenConsumerTridentTokenExchangedForZetaIterator, error) {

	logs, sub, err := _ZetaTokenConsumerTrident.contract.FilterLogs(opts, "TokenExchangedForZeta")
	if err != nil {
		return nil, err
	}
	return &ZetaTokenConsumerTridentTokenExchangedForZetaIterator{contract: _ZetaTokenConsumerTrident.contract, event: "TokenExchangedForZeta", logs: logs, sub: sub}, nil
}

// WatchTokenExchangedForZeta is a free log subscription operation binding the contract event 0x017190d3d99ee6d8dd0604ef1e71ff9802810c6485f57c9b2ed6169848dd119f.
//
// Solidity: event TokenExchangedForZeta(address token, uint256 amountIn, uint256 amountOut)
func (_ZetaTokenConsumerTrident *ZetaTokenConsumerTridentFilterer) WatchTokenExchangedForZeta(opts *bind.WatchOpts, sink chan<- *ZetaTokenConsumerTridentTokenExchangedForZeta) (event.Subscription, error) {

	logs, sub, err := _ZetaTokenConsumerTrident.contract.WatchLogs(opts, "TokenExchangedForZeta")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaTokenConsumerTridentTokenExchangedForZeta)
				if err := _ZetaTokenConsumerTrident.contract.UnpackLog(event, "TokenExchangedForZeta", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTokenExchangedForZeta is a log parse operation binding the contract event 0x017190d3d99ee6d8dd0604ef1e71ff9802810c6485f57c9b2ed6169848dd119f.
//
// Solidity: event TokenExchangedForZeta(address token, uint256 amountIn, uint256 amountOut)
func (_ZetaTokenConsumerTrident *ZetaTokenConsumerTridentFilterer) ParseTokenExchangedForZeta(log types.Log) (*ZetaTokenConsumerTridentTokenExchangedForZeta, error) {
	event := new(ZetaTokenConsumerTridentTokenExchangedForZeta)
	if err := _ZetaTokenConsumerTrident.contract.UnpackLog(event, "TokenExchangedForZeta", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaTokenConsumerTridentZetaExchangedForEthIterator is returned from FilterZetaExchangedForEth and is used to iterate over the raw logs and unpacked data for ZetaExchangedForEth events raised by the ZetaTokenConsumerTrident contract.
type ZetaTokenConsumerTridentZetaExchangedForEthIterator struct {
	Event *ZetaTokenConsumerTridentZetaExchangedForEth // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZetaTokenConsumerTridentZetaExchangedForEthIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaTokenConsumerTridentZetaExchangedForEth)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ZetaTokenConsumerTridentZetaExchangedForEth)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ZetaTokenConsumerTridentZetaExchangedForEthIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaTokenConsumerTridentZetaExchangedForEthIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaTokenConsumerTridentZetaExchangedForEth represents a ZetaExchangedForEth event raised by the ZetaTokenConsumerTrident contract.
type ZetaTokenConsumerTridentZetaExchangedForEth struct {
	AmountIn  *big.Int
	AmountOut *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterZetaExchangedForEth is a free log retrieval operation binding the contract event 0x74e171117e91660f493740924d8bad0caf48dc4fbccb767fb05935397a2c17ae.
//
// Solidity: event ZetaExchangedForEth(uint256 amountIn, uint256 amountOut)
func (_ZetaTokenConsumerTrident *ZetaTokenConsumerTridentFilterer) FilterZetaExchangedForEth(opts *bind.FilterOpts) (*ZetaTokenConsumerTridentZetaExchangedForEthIterator, error) {

	logs, sub, err := _ZetaTokenConsumerTrident.contract.FilterLogs(opts, "ZetaExchangedForEth")
	if err != nil {
		return nil, err
	}
	return &ZetaTokenConsumerTridentZetaExchangedForEthIterator{contract: _ZetaTokenConsumerTrident.contract, event: "ZetaExchangedForEth", logs: logs, sub: sub}, nil
}

// WatchZetaExchangedForEth is a free log subscription operation binding the contract event 0x74e171117e91660f493740924d8bad0caf48dc4fbccb767fb05935397a2c17ae.
//
// Solidity: event ZetaExchangedForEth(uint256 amountIn, uint256 amountOut)
func (_ZetaTokenConsumerTrident *ZetaTokenConsumerTridentFilterer) WatchZetaExchangedForEth(opts *bind.WatchOpts, sink chan<- *ZetaTokenConsumerTridentZetaExchangedForEth) (event.Subscription, error) {

	logs, sub, err := _ZetaTokenConsumerTrident.contract.WatchLogs(opts, "ZetaExchangedForEth")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaTokenConsumerTridentZetaExchangedForEth)
				if err := _ZetaTokenConsumerTrident.contract.UnpackLog(event, "ZetaExchangedForEth", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseZetaExchangedForEth is a log parse operation binding the contract event 0x74e171117e91660f493740924d8bad0caf48dc4fbccb767fb05935397a2c17ae.
//
// Solidity: event ZetaExchangedForEth(uint256 amountIn, uint256 amountOut)
func (_ZetaTokenConsumerTrident *ZetaTokenConsumerTridentFilterer) ParseZetaExchangedForEth(log types.Log) (*ZetaTokenConsumerTridentZetaExchangedForEth, error) {
	event := new(ZetaTokenConsumerTridentZetaExchangedForEth)
	if err := _ZetaTokenConsumerTrident.contract.UnpackLog(event, "ZetaExchangedForEth", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaTokenConsumerTridentZetaExchangedForTokenIterator is returned from FilterZetaExchangedForToken and is used to iterate over the raw logs and unpacked data for ZetaExchangedForToken events raised by the ZetaTokenConsumerTrident contract.
type ZetaTokenConsumerTridentZetaExchangedForTokenIterator struct {
	Event *ZetaTokenConsumerTridentZetaExchangedForToken // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZetaTokenConsumerTridentZetaExchangedForTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaTokenConsumerTridentZetaExchangedForToken)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ZetaTokenConsumerTridentZetaExchangedForToken)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ZetaTokenConsumerTridentZetaExchangedForTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaTokenConsumerTridentZetaExchangedForTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaTokenConsumerTridentZetaExchangedForToken represents a ZetaExchangedForToken event raised by the ZetaTokenConsumerTrident contract.
type ZetaTokenConsumerTridentZetaExchangedForToken struct {
	Token     common.Address
	AmountIn  *big.Int
	AmountOut *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterZetaExchangedForToken is a free log retrieval operation binding the contract event 0x0a7cb8f6e1d29e616c1209a3f418c17b3a9137005377f6dd072217b1ede2573b.
//
// Solidity: event ZetaExchangedForToken(address token, uint256 amountIn, uint256 amountOut)
func (_ZetaTokenConsumerTrident *ZetaTokenConsumerTridentFilterer) FilterZetaExchangedForToken(opts *bind.FilterOpts) (*ZetaTokenConsumerTridentZetaExchangedForTokenIterator, error) {

	logs, sub, err := _ZetaTokenConsumerTrident.contract.FilterLogs(opts, "ZetaExchangedForToken")
	if err != nil {
		return nil, err
	}
	return &ZetaTokenConsumerTridentZetaExchangedForTokenIterator{contract: _ZetaTokenConsumerTrident.contract, event: "ZetaExchangedForToken", logs: logs, sub: sub}, nil
}

// WatchZetaExchangedForToken is a free log subscription operation binding the contract event 0x0a7cb8f6e1d29e616c1209a3f418c17b3a9137005377f6dd072217b1ede2573b.
//
// Solidity: event ZetaExchangedForToken(address token, uint256 amountIn, uint256 amountOut)
func (_ZetaTokenConsumerTrident *ZetaTokenConsumerTridentFilterer) WatchZetaExchangedForToken(opts *bind.WatchOpts, sink chan<- *ZetaTokenConsumerTridentZetaExchangedForToken) (event.Subscription, error) {

	logs, sub, err := _ZetaTokenConsumerTrident.contract.WatchLogs(opts, "ZetaExchangedForToken")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaTokenConsumerTridentZetaExchangedForToken)
				if err := _ZetaTokenConsumerTrident.contract.UnpackLog(event, "ZetaExchangedForToken", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseZetaExchangedForToken is a log parse operation binding the contract event 0x0a7cb8f6e1d29e616c1209a3f418c17b3a9137005377f6dd072217b1ede2573b.
//
// Solidity: event ZetaExchangedForToken(address token, uint256 amountIn, uint256 amountOut)
func (_ZetaTokenConsumerTrident *ZetaTokenConsumerTridentFilterer) ParseZetaExchangedForToken(log types.Log) (*ZetaTokenConsumerTridentZetaExchangedForToken, error) {
	event := new(ZetaTokenConsumerTridentZetaExchangedForToken)
	if err := _ZetaTokenConsumerTrident.contract.UnpackLog(event, "ZetaExchangedForToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
