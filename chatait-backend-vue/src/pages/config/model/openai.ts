/*
 * Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
 * Use of this source code is governed by a AGPL v3.0 style
 * license that can be found in the LICENSE file.
 */

export interface FormOpenaiEdit {
  id: string
  title: string
  api_url: string
  api_key: string
  proxy: string
  max_tokens: number
  gpt3_model: string
  gpt4_model: string
  status: number
}
