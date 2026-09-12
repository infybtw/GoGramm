package api

import (
	"encoding/json"
	"testing"
)

func TestDecodeCoreUnionTypes(t *testing.T) {
	t.Run("callback query message", func(t *testing.T) {
		var update Update
		if err := json.Unmarshal([]byte(`{"update_id":1,"callback_query":{"id":"id","from":{},"chat_instance":"chat","message":{"message_id":1,"date":1,"chat":{"id":1,"type":"private"}}}}`), &update); err != nil {
			t.Fatalf("decode update: %v", err)
		}
		if _, ok := update.CallbackQuery.Message.(*Message); !ok {
			t.Fatalf("message = %T, want *Message", update.CallbackQuery.Message)
		}
	})

	t.Run("chat member update", func(t *testing.T) {
		var update Update
		if err := json.Unmarshal([]byte(`{"update_id":1,"my_chat_member":{"chat":{"id":1,"type":"private"},"from":{},"date":1,"old_chat_member":{"status":"member","user":{}},"new_chat_member":{"status":"administrator","user":{}}}}`), &update); err != nil {
			t.Fatalf("decode update: %v", err)
		}
		if _, ok := update.MyChatMember.OldChatMember.(*ChatMemberMember); !ok {
			t.Fatalf("old_chat_member = %T, want *ChatMemberMember", update.MyChatMember.OldChatMember)
		}
		if _, ok := update.MyChatMember.NewChatMember.(*ChatMemberAdministrator); !ok {
			t.Fatalf("new_chat_member = %T, want *ChatMemberAdministrator", update.MyChatMember.NewChatMember)
		}
	})

	t.Run("paid media", func(t *testing.T) {
		var update Update
		if err := json.Unmarshal([]byte(`{"update_id":1,"message":{"message_id":1,"date":1,"chat":{"id":1,"type":"private"},"paid_media":{"star_count":2,"paid_media":[{"type":"preview","width":100}]}}}`), &update); err != nil {
			t.Fatalf("decode update: %v", err)
		}
		if _, ok := update.Message.PaidMedia.PaidMedia[0].(*PaidMediaPreview); !ok {
			t.Fatalf("paid media = %T, want *PaidMediaPreview", update.Message.PaidMedia.PaidMedia[0])
		}
	})

	t.Run("reactions", func(t *testing.T) {
		var update Update
		if err := json.Unmarshal([]byte(`{"update_id":1,"message_reaction":{"chat":{"id":1,"type":"private"},"message_id":1,"date":1,"old_reaction":[{"type":"emoji","emoji":"x"}],"new_reaction":[{"type":"paid"}]}}`), &update); err != nil {
			t.Fatalf("decode update: %v", err)
		}
		if _, ok := update.MessageReaction.OldReaction[0].(*ReactionTypeEmoji); !ok {
			t.Fatalf("old reaction = %T, want *ReactionTypeEmoji", update.MessageReaction.OldReaction[0])
		}
		if _, ok := update.MessageReaction.NewReaction[0].(*ReactionTypePaid); !ok {
			t.Fatalf("new reaction = %T, want *ReactionTypePaid", update.MessageReaction.NewReaction[0])
		}
		if err := json.Unmarshal([]byte(`{"update_id":1,"message_reaction_count":{"chat":{"id":1,"type":"private"},"message_id":1,"date":1,"reactions":[{"type":{"type":"custom_emoji","custom_emoji_id":"id"},"total_count":1}]}}`), &update); err != nil {
			t.Fatalf("decode reaction count update: %v", err)
		}
		if _, ok := update.MessageReactionCount.Reactions[0].Type.(*ReactionTypeCustomEmoji); !ok {
			t.Fatalf("reaction count type = %T, want *ReactionTypeCustomEmoji", update.MessageReactionCount.Reactions[0].Type)
		}
	})
}

func TestDecodeGiftsAndStarTransactions(t *testing.T) {
	var gifts OwnedGifts
	if err := json.Unmarshal([]byte(`{"total_count":2,"gifts":[{"type":"regular","gift":{}},{"type":"unique","gift":{}}]}`), &gifts); err != nil {
		t.Fatalf("decode owned gifts: %v", err)
	}
	if _, ok := gifts.Gifts[0].(*OwnedGiftRegular); !ok {
		t.Fatalf("first gift = %T, want *OwnedGiftRegular", gifts.Gifts[0])
	}
	if _, ok := gifts.Gifts[1].(*OwnedGiftUnique); !ok {
		t.Fatalf("second gift = %T, want *OwnedGiftUnique", gifts.Gifts[1])
	}

	var transactions StarTransactions
	if err := json.Unmarshal([]byte(`{"transactions":[{"id":"id","amount":1,"date":1,"source":{"type":"user","user":{},"paid_media":[{"type":"photo","photo":[]}]},"receiver":{"type":"fragment","withdrawal_state":{"type":"succeeded","date":1,"url":"https://example.com"}}}]}`), &transactions); err != nil {
		t.Fatalf("decode star transactions: %v", err)
	}
	transaction := transactions.Transactions[0]
	if _, ok := transaction.Source.(*TransactionPartnerUser); !ok {
		t.Fatalf("source = %T, want *TransactionPartnerUser", transaction.Source)
	}
	fragment, ok := transaction.Receiver.(*TransactionPartnerFragment)
	if !ok {
		t.Fatalf("receiver = %T, want *TransactionPartnerFragment", transaction.Receiver)
	}
	if _, ok := fragment.WithdrawalState.(*RevenueWithdrawalStateSucceeded); !ok {
		t.Fatalf("withdrawal state = %T, want *RevenueWithdrawalStateSucceeded", fragment.WithdrawalState)
	}

	if err := json.Unmarshal([]byte(`{"transactions":[{"id":"id","amount":1,"date":1}]}`), &transactions); err != nil {
		t.Fatalf("decode transaction without partners: %v", err)
	}
	if transactions.Transactions[0].Source != nil || transactions.Transactions[0].Receiver != nil {
		t.Fatalf("transaction partners = %T/%T, want nil/nil", transactions.Transactions[0].Source, transactions.Transactions[0].Receiver)
	}
}

func TestBotCommandScopeChatIDSupportsUsername(t *testing.T) {
	b, err := json.Marshal(&BotCommandScopeChat{ChatID: NewUsername("@channel")})
	if err != nil {
		t.Fatalf("marshal scope: %v", err)
	}
	if string(b) != `{"chat_id":"@channel","type":"chat"}` {
		t.Fatalf("scope JSON = %s", b)
	}
}
