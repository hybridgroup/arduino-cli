/*
 * This file is part of arduino-cli.
 *
 * arduino-cli is free software; you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation; either version 2 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program; if not, write to the Free Software
 * Foundation, Inc., 51 Franklin St, Fifth Floor, Boston, MA  02110-1301  USA
 *
 * As a special exception, you may use this file as part of a free software
 * library without restriction.  Specifically, if other files instantiate
 * templates or use macros or inline functions from this file, or you compile
 * this file and link it with other files to produce an executable, this
 * file does not by itself cause the resulting executable to be covered by
 * the GNU General Public License.  This exception does not however
 * invalidate any other reasons why the executable file might be covered by
 * the GNU General Public License.
 *
 * Copyright 2017 BCMI LABS SA (http://www.arduino.cc/)
 */

package builderClient

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

// ListBoardsPath computes a request path to the list action of boards.
func ListBoardsPath() string {
	return fmt.Sprintf("/builder/v1/boards")
}

// ListBoards provides a list of all the boards supported by Arduino Create. Doesn't require any authentication.
func (c *Client) ListBoards(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewListBoardsRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewListBoardsRequest create the request corresponding to the list action endpoint of the boards resource.
func (c *Client) NewListBoardsRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// ShowBoardsPath computes a request path to the show action of boards.
func ShowBoardsPath(vid string, pid string) string {
	param0 := vid
	param1 := pid

	return fmt.Sprintf("/builder/v1/boards/%s/%s", param0, param1)
}

// ShowBoards provides the board identified by the :vid and :pid params. Doesn't require authentication.
func (c *Client) ShowBoards(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewShowBoardsRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewShowBoardsRequest create the request corresponding to the show action endpoint of the boards resource.
func (c *Client) NewShowBoardsRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}