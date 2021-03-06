// Copyright 2020 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package post

import (
	"github.com/gin-gonic/gin"

	"github.com/marmotedu/goserver/internal/pkg/constant"
	"github.com/marmotedu/goserver/internal/pkg/log"
	"github.com/marmotedu/goserver/pkg/core"
	metav1 "github.com/marmotedu/goserver/pkg/meta/v1"
)

// Get get a post by the post identifier.
func (p *PostController) Get(c *gin.Context) {
	log.L(c).Info("get post function called.")

	post, err := p.srv.Posts().Get(c, c.GetString(constant.XUsernameKey), c.Param("postID"), metav1.GetOptions{})
	if err != nil {
		core.WriteResponse(c, err, nil)

		return
	}

	core.WriteResponse(c, nil, post)
}
