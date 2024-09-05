/*
 * Licensed to the OpenAirInterface (OAI) Software Alliance under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The OpenAirInterface Software Alliance licenses this file to You under
 * the terms found in the LICENSE file in the root of this source tree.
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *-------------------------------------------------------------------------------
 * For more information about the OpenAirInterface (OAI) Software Alliance:
 *      contact@openairinterface.org
 */

#include "lte/gateway/c/core/oai/tasks/nas/ies/LinkedEpsBearerIdentity.hpp"

#include <stdint.h>

#ifdef __cplusplus
extern "C" {
#endif
#include "lte/gateway/c/core/oai/common/TLVEncoder.h"
#include "lte/gateway/c/core/oai/common/TLVDecoder.h"
#ifdef __cplusplus
}
#endif

//------------------------------------------------------------------------------
int decode_linked_eps_bearer_identity(
    linked_eps_bearer_identity_t* linkedepsbeareridentity, uint8_t iei,
    uint8_t* buffer, uint32_t len) {
  int decoded = 0;

  CHECK_PDU_POINTER_AND_LENGTH_DECODER(
      buffer, LINKED_EPS_BEARER_IDENTITY_MINIMUM_LENGTH, len);

  if (iei > 0) {
    CHECK_IEI_DECODER((*buffer & 0xf0), iei);
  }

  *linkedepsbeareridentity = *buffer & 0xf;
  decoded++;
  return decoded;
}

//------------------------------------------------------------------------------
int decode_u8_linked_eps_bearer_identity(
    linked_eps_bearer_identity_t* linkedepsbeareridentity, uint8_t iei,
    uint8_t value, uint32_t len) {
  int decoded = 0;
  uint8_t* buffer = &value;

  *linkedepsbeareridentity = *buffer & 0xf;
  decoded++;
  return decoded;
}

//------------------------------------------------------------------------------
int encode_linked_eps_bearer_identity(
    linked_eps_bearer_identity_t* linkedepsbeareridentity, uint8_t iei,
    uint8_t* buffer, uint32_t len) {
  uint8_t encoded = 0;

  /*
   * Checking length and pointer
   */
  CHECK_PDU_POINTER_AND_LENGTH_ENCODER(
      buffer, LINKED_EPS_BEARER_IDENTITY_MINIMUM_LENGTH, len);
  *(buffer + encoded) = 0x00 | (iei & 0xf0) | (*linkedepsbeareridentity & 0xf);
  encoded++;
  return encoded;
}

//------------------------------------------------------------------------------
uint8_t encode_u8_linked_eps_bearer_identity(
    linked_eps_bearer_identity_t* linkedepsbeareridentity) {
  uint8_t bufferReturn;
  uint8_t* buffer = &bufferReturn;
  uint8_t encoded = 0;
  uint8_t iei = 0;

  *(buffer + encoded) = 0x00 | (iei & 0xf0) | (*linkedepsbeareridentity & 0xf);
  encoded++;
  return bufferReturn;
}
