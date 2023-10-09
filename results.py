# %%
import pandas as pd


# %%
generic_runs = pd. read_csv('./runs/generic_runs.csv',header=None,names=['generic_Runtime'])
nongeneric_runs = pd. read_csv('./runs/nongeneric_runs.csv',header=None,names=['nongeneric_Runtime'])

# %%
generic_runs_split = pd.DataFrame({'generic_int_runtime':generic_runs['generic_Runtime'].iloc[::2].values, 
                                   'generic_float_runtime':generic_runs['generic_Runtime'].iloc[1::2].values})

nongeneric_runs_split = pd.DataFrame({'nongeneric_int_runtime':nongeneric_runs['nongeneric_Runtime'].iloc[::2].values, 
                                   'nongeneric_float_runtime':nongeneric_runs['nongeneric_Runtime'].iloc[1::2].values})

merged_df = pd.merge(generic_runs_split, nongeneric_runs_split, left_index=True, right_index=True)
#print(generic_runs_split.head(1))
# %%
print(merged_df.describe())


