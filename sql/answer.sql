-- Soal Nomor 1

select * from t_policy tp 
join t_client tc on tp.client_number = tc.client_number
where tp.policy_submit_date > '2018-01-15' and extract(month from tc.birth_date) = 9;

-- Soal Nomor 2

select * from t_policy tp 
join t_agent ta on tp.agent_code = ta.agent_code 
where tp.policy_status = 'INFORCE' and ta.agent_office = 'JAKARTA';

-- Soal Nomor 3

select ta.agent_code, ta.agent_name, tp.commission / tp.premium * 100 as basic_commission
from t_agent ta 
join t_policy tp on ta.agent_code = tp.agent_code 

-- Soal Nomor 4
-- Kurang tau saya caranya gimana

-- Soal Nomor 5

select tp.policy_number, ta.agent_code, tp.premium, tp.discount, tp.premium - (tp.premium * tp.discount / 100) as premium_after_discount
from t_policy tp 
join t_agent ta on tp.agent_code = ta.agent_code 
where tp.premium - (tp.premium * tp.discount / 100) < 1000000
order by premium_after_discount asc;

