// Package logic
// Written by yijian on 2024/01/20
package logic

import (
	"github.com/zeromicro/go-zero/core/logc"
	"mooon-mailbox/model"
	mooonmailbox "mooon-mailbox/pb/mooon-mailbox"
	"time"
)

func deliverMessage(l *DeliverMessageLogic, in *mooonmailbox.DeliverMessageReq) (*mooonmailbox.DeliverMessageResp, error) {
	// todo: add your logic here and delete this line
	now := time.Now()
	letter := model.TMooonMailbox{
		FRecipient:   in.Recipient,
		FDeliverTime: now,
		FArrivalTime: now,
		FState:       int64(mooonmailbox.ListMessagesReq_LA_ONLY_UNREAD),
		FLetterBody:  in.LetterBody,
	}
	dbResult, err := l.svcCtx.MailboxModel.Insert(l.ctx, &letter)
	if err != nil {
		logc.Errorf(l.ctx, "Insert %s to table error: %s\n", in.String(), err.Error())
		return nil, err
	} else {
		var rowsAffected int64 = -1
		var lastInsertId int64 = -1
		rowsAffected, _ = dbResult.RowsAffected()
		lastInsertId, _ = dbResult.LastInsertId()
		logc.Infof(l.ctx, "Insert %s success: RowsAffected=%d, LastInsertId=%d\n", in.String(), rowsAffected, lastInsertId)
		return &mooonmailbox.DeliverMessageResp{
			Recipient: in.Recipient,
		}, nil
	}
}