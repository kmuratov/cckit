// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: erc20_utxo/service/allowance/allowance.proto

package allowance

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/hyperledger-labs/cckit/extensions/token"
	_ "github.com/mwitkow/go-proto-validators"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	_ "google.golang.org/protobuf/types/known/emptypb"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *ApproveRequest) Validate() error {
	if this.Owner == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Owner", fmt.Errorf(`value '%v' must not be an empty string`, this.Owner))
	}
	if this.Spender == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Spender", fmt.Errorf(`value '%v' must not be an empty string`, this.Spender))
	}
	if this.Symbol == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Symbol", fmt.Errorf(`value '%v' must not be an empty string`, this.Symbol))
	}
	if this.Amount != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Amount); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Amount", err)
		}
	}
	return nil
}
func (this *TransferFromRequest) Validate() error {
	if this.Owner == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Owner", fmt.Errorf(`value '%v' must not be an empty string`, this.Owner))
	}
	if this.Recipient == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Recipient", fmt.Errorf(`value '%v' must not be an empty string`, this.Recipient))
	}
	if this.Symbol == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Symbol", fmt.Errorf(`value '%v' must not be an empty string`, this.Symbol))
	}
	if this.Amount != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Amount); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Amount", err)
		}
	}
	return nil
}
func (this *TransferFromResponse) Validate() error {
	if this.Amount != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Amount); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Amount", err)
		}
	}
	return nil
}
func (this *AllowanceId) Validate() error {
	if this.Owner == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Owner", fmt.Errorf(`value '%v' must not be an empty string`, this.Owner))
	}
	if this.Spender == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Spender", fmt.Errorf(`value '%v' must not be an empty string`, this.Spender))
	}
	if this.Symbol == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Symbol", fmt.Errorf(`value '%v' must not be an empty string`, this.Symbol))
	}
	return nil
}
func (this *Allowance) Validate() error {
	if this.Amount != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Amount); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Amount", err)
		}
	}
	return nil
}
func (this *Operation) Validate() error {
	if this.Amount != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Amount); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Amount", err)
		}
	}
	return nil
}
func (this *Allowances) Validate() error {
	for _, item := range this.Items {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Items", err)
			}
		}
	}
	return nil
}
func (this *Approved) Validate() error {
	if this.Amount != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Amount); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Amount", err)
		}
	}
	return nil
}
func (this *TransferredFrom) Validate() error {
	if this.Amount != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Amount); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Amount", err)
		}
	}
	return nil
}
