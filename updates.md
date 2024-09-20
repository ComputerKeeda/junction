

ignite scaffold type final_espresso_state state nounce:uint

change the proto to byte from string

ignite scaffold message set_espresso_finalized_state finalized_state nounce:uint --response status:bool

also change its finalizedState type from string to byte

ignite scaffold query get_espresso_state nounce:uint --response espresso_state:FinalEspressoState