// Copyright Weidong Liang 2013. All rights reserved.

package main

import (
	"testing"
)

func Test_StripCodeSectionFromHTML(t *testing.T) {
	test_str := "<p>I am import matlab file and construct a data frame, matlab file" +
		"contains two columns with and each row maintain a cell that has a matrix, I" +
		" construct a dataframe to run random forest. But I am getting following error." +
		" </p>\n\n<pre><code>Error in model.frame.default(formula = expert_data_frame$t_labels ~ .,  : " +
		" invalid type (list) for variable 'expert_data_frame$t_labels'" +
		" </code></pre>\n<p>Here is the code how I import the matlab file and construct" +
		" the dataframe:</p>\n\n<pre><code>all_exp_traintest &lt;- readMat(all_exp_filepath);" +
		" len = length(all_exp_traintest$exp.traintest)/2;\n for (i in 1:len) {\n" +
		" expert_train_df &lt;- data.frame(all_exp_traintest$exp.traintest[i]);" +
		" labels = data.frame(all_exp_traintest$exp.traintest[i+302]);" +
		" names(labels)[1] &lt;- \"\"t_labels\"\";" +
		"       expert_train_df$t_labels &lt;- labels;" +
		"       expert_data_frame &lt;- data.frame(expert_train_df);" +
		"     rf_model = randomForest(expert_data_frame$t_labels ~., data=expert_data_frame, importance=TRUE, do.trace=100);" +
		" }" +
		" </code></pre>" +
		" <p>Structure of the Matlab input file</p>"
	test_result := "I am import matlab file and construct a data frame, matlab file" +
		"contains two columns with and each row maintain a cell that has a matrix, I" +
		" construct a dataframe to run random forest. But I am getting following error.    " +
		"Here is the code how I import the matlab file and construct" +
		" the dataframe:  " +
		" Structure of the Matlab input file"
	r := StripCodeSectionFromHTML(test_str)
	if r != test_result {
		t.Errorf("Expected \n%s\n but got\n%s\n", test_result, r)
	}
}
