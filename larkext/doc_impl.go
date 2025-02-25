/**
 * Copyright 2022 chyroc
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package larkext

import (
	"context"
	"encoding/json"

	"github.com/chyroc/lark"
)

func (r *Doc) meta(ctx context.Context) (*lark.GetDriveDocMetaResp, error) {
	resp, _, err := r.larkClient.Drive.GetDriveDocMeta(ctx, &lark.GetDriveDocMetaReq{
		DocToken: r.docToken,
	})
	return resp, err
}

func (r *Doc) delete(ctx context.Context) error {
	_, _, err := r.larkClient.Drive.DeleteDriveFile(ctx, &lark.DeleteDriveFileReq{
		Type:      "doc",
		FileToken: r.docToken,
	})
	return err
}

func (r *Doc) rawContent(ctx context.Context) (string, error) {
	resp, _, err := r.larkClient.Drive.GetDriveDocRawContent(ctx, &lark.GetDriveDocRawContentReq{
		DocToken: r.docToken,
	})
	return resp.Content, err
}

func (r *Doc) content(ctx context.Context) (*lark.DocContent, error) {
	resp, _, err := r.larkClient.Drive.GetDriveDocContent(ctx, &lark.GetDriveDocContentReq{
		DocToken: r.docToken,
	})
	doc := &lark.DocContent{}
	err = json.Unmarshal([]byte(resp.Content), doc)
	if err != nil {
		return nil, err
	}
	return doc, err
}

func (r *Doc) update(ctx context.Context, requests ...*lark.UpdateDocRequest) error {
	revision, err := r.revision(ctx)
	if err != nil {
		return err
	}
	_, _, err = r.larkClient.Drive.UpdateDriveDocContent(ctx, &lark.UpdateDriveDocContentReq{
		DocToken: r.docToken,
		Revision: revision,
		Requests: lark.UpdateDocRequests(requests).ToString(),
	})
	return err
}

func (r *Doc) copy(ctx context.Context, folderToken, name string) (*Doc, error) {
	res, err := copyFile(ctx, r.larkClient, folderToken, r.docToken, r.typ, name)
	if err != nil {
		return nil, err
	}
	return newDoc(r.larkClient, res.Token, res.URL), nil
}

func (r *Doc) revision(ctx context.Context) (int64, error) {
	resp, _, err := r.larkClient.Drive.GetDriveDocContent(ctx, &lark.GetDriveDocContentReq{
		DocToken: r.docToken,
	})
	if err != nil {
		return 0, err
	}
	return resp.Revision, nil
}

func (r *Doc) uploadMedia(ctx context.Context, parentType, parentToken string) {
	r.larkClient.Drive.UploadDriveMedia(ctx, &lark.UploadDriveMediaReq{
		FileName:   "",
		ParentType: parentType,
		ParentNode: parentToken,
		Size:       0,
		Checksum:   nil,
		Extra:      nil,
		File:       nil,
	})
}
