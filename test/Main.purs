module Test.Main where

import Prelude
import Effect (Effect)
import Effect.Console (log)
import JS.BigInt (fromInt, toString)

main :: Effect Unit
main = do
  log "Running BigInt tests..."
  let a = fromInt 10
  let b = fromInt 20
  let c = a + b
  if c == fromInt 30 then
    log "10 + 20 = 30 (OK)"
  else
    log $ "Failed: expected 30, got " <> toString c
