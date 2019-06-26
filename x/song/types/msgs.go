package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const RouterKey = ModuleName // this was defined in your key.go file

// MsgPublish defines a Publish message
type MsgPublish struct {
	Title string         `json:"title"`
	Owner sdk.AccAddress `json:"owner"`
}

// NewMsgPublish is a constructor function for MsgPublish
func NewMsgPublish(title string, owner sdk.AccAddress) MsgPublish {
	return MsgPublish{
		Title: title,
		Owner: owner,
	}
}

// ValidateBasic runs stateless checks on the message
func (msg MsgPublish) ValidateBasic() sdk.Error {
	if msg.Owner.Empty() {
		return sdk.ErrInvalidAddress(msg.Owner.String())
	}
	if len(msg.Title) == 0 {
		return sdk.ErrUnknownRequest("Title cannot be empty")
	}

	return nil
}

// Route should return the name of the module
func (msg MsgPublish) Route() string { return RouterKey }

// Type should return the action
func (msg MsgPublish) Type() string { return "publish" }

// GetSignBytes encodes the message for signing
func (msg MsgPublish) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgPublish) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Owner}
}
