using System;
using System.Data;
using System.Data.SQLite;

namespace testsqlite
{
    class Program
    {
        //static readonly string ConnectionString = "Data Source=:memory:";
        static readonly string ConnectionString = "Data Source=test.db;Version=3;";

        static void Main(string[] args)
        {
            using SQLiteConnection conn = new SQLiteConnection(ConnectionString);
            conn.Open();

            ExecPramgas(conn);
            CreateTable(conn);
            //InsertRows(conn);
            InsertRows2(conn);
        }

        static void CreateTable(SQLiteConnection conn)
        {
            using var cmd = new SQLiteCommand();
            cmd.Connection = conn;

            cmd.CommandText = @"CREATE TABLE IF NOT EXISTS userinfo 
		        (
                    id INTEGER not null primary key
                    ,username CHAR(10)
                    ,departname char(10) not null
                    ,created date not null
		        )";

            try
            {
                cmd.ExecuteNonQuery();
            }
            catch (SQLiteException e)
            {
                Console.WriteLine(e.Message);
            }
        }

        static void ExecPramgas(SQLiteConnection conn)
        {
            using SQLiteCommand cmd = new SQLiteCommand();
            cmd.Connection = conn;

            try
            {
                cmd.CommandText = "PRAGMA journal_mode = OFF;";
                cmd.ExecuteNonQuery();

                cmd.CommandText = "PRAGMA synchronous = OFF;";
                cmd.ExecuteNonQuery();

                cmd.CommandText = "PRAGMA temp_store = MEMORY;";
                cmd.ExecuteNonQuery();
            }
            catch (SQLiteException e)
            {
                Console.WriteLine(e.Message);
            }
        }

        static void InsertRows(SQLiteConnection conn)
        {
            using var cmd = new SQLiteCommand();
            cmd.Connection = conn;

            cmd.CommandText = "BEGIN";
            cmd.ExecuteNonQuery();

            for (int i = 0; i < 2_000_000; i++)
            {
                cmd.CommandText = @"INSERT INTO userinfo (username,departname,created) 
                VALUES (@username,@departname,@created),
                    (@username,@departname,@created),
                    (@username,@departname,@created),
                    (@username,@departname,@created),
                    (@username,@departname,@created),
                    (@username,@departname,@created),
                    (@username,@departname,@created),
                    (@username,@departname,@created),
                    (@username,@departname,@created),
                    (@username,@departname,@created),
                    (@username,@departname,@created),
                    (@username,@departname,@created),
                    (@username,@departname,@created),
                    (@username,@departname,@created),
                    (@username,@departname,@created),
                    (@username,@departname,@created),
                    (@username,@departname,@created),
                    (@username,@departname,@created),
                    (@username,@departname,@created),
                    (@username,@departname,@created),
                    (@username,@departname,@created),
                    (@username,@departname,@created),
                    (@username,@departname,@created),
                    (@username,@departname,@created),
                    (@username,@departname,@created),                    
                    (@username,@departname,@created),
                    (@username,@departname,@created),
                    (@username,@departname,@created),
                    (@username,@departname,@created),
                    (@username,@departname,@created),
                    (@username,@departname,@created),
                    (@username,@departname,@created),
                    (@username,@departname,@created),
                    (@username,@departname,@created),
                    (@username,@departname,@created),
                    (@username,@departname,@created),
                    (@username,@departname,@created),
                    (@username,@departname,@created),
                    (@username,@departname,@created),
                    (@username,@departname,@created),
                    (@username,@departname,@created),
                    (@username,@departname,@created),
                    (@username,@departname,@created),
                    (@username,@departname,@created),
                    (@username,@departname,@created),
                    (@username,@departname,@created),
                    (@username,@departname,@created),
                    (@username,@departname,@created),
                    (@username,@departname,@created),
                    
                    (@username,@departname,@created)                 
                ;";
                
                cmd.Parameters.AddWithValue("@username", "zhangsan");
                cmd.Parameters.AddWithValue("@departname", "IT");
                cmd.Parameters.AddWithValue("@created", DateTime.Now);
                cmd.ExecuteNonQuery();
            }

            cmd.CommandText = "COMMIT";
            cmd.ExecuteNonQuery();
        }
 
        static void InsertRows2(SQLiteConnection conn)
        {
            SQLiteTransaction tx = conn.BeginTransaction();
            try
            {
                SQLiteCommand cmd = new SQLiteCommand();
                cmd.Transaction = tx;
                cmd.CommandType = CommandType.Text;
                cmd.CommandText = @"INSERT INTO userinfo (username,departname,created) 
                VALUES (@username,@departname,@created),
                    (@username,@departname,@created),
                    (@username,@departname,@created),
                    (@username,@departname,@created),
                    (@username,@departname,@created),
                    (@username,@departname,@created),
                    (@username,@departname,@created),
                    (@username,@departname,@created),
                    (@username,@departname,@created),
                    (@username,@departname,@created),
                    (@username,@departname,@created),
                    (@username,@departname,@created),
                    (@username,@departname,@created),
                    (@username,@departname,@created),
                    (@username,@departname,@created),
                    (@username,@departname,@created),
                    (@username,@departname,@created),
                    (@username,@departname,@created),
                    (@username,@departname,@created),
                    (@username,@departname,@created),
                    (@username,@departname,@created),
                    (@username,@departname,@created),
                    (@username,@departname,@created),
                    (@username,@departname,@created),
                    (@username,@departname,@created),                    
                    (@username,@departname,@created),
                    (@username,@departname,@created),
                    (@username,@departname,@created),
                    (@username,@departname,@created),
                    (@username,@departname,@created),
                    (@username,@departname,@created),
                    (@username,@departname,@created),
                    (@username,@departname,@created),
                    (@username,@departname,@created),
                    (@username,@departname,@created),
                    (@username,@departname,@created),
                    (@username,@departname,@created),
                    (@username,@departname,@created),
                    (@username,@departname,@created),
                    (@username,@departname,@created),
                    (@username,@departname,@created),
                    (@username,@departname,@created),
                    (@username,@departname,@created),
                    (@username,@departname,@created),
                    (@username,@departname,@created),
                    (@username,@departname,@created),
                    (@username,@departname,@created),
                    (@username,@departname,@created),
                    (@username,@departname,@created),
                    
                    (@username,@departname,@created)                 
                ;";
                cmd.Parameters.Add(new SQLiteParameter("@username", DbType.String));
                cmd.Parameters.Add(new SQLiteParameter("@departname", DbType.String));
                cmd.Parameters.Add(new SQLiteParameter("@created", DbType.DateTime));

                DateTime dt = DateTime.Now;
                for (int i = 0; i < 2_000_000; i++)
                {
                    cmd.Parameters[0].Value = "zhangsan";
                    cmd.Parameters[1].Value = "IT";
                    cmd.Parameters[2].Value = dt;
                    cmd.ExecuteNonQuery();
                }
                tx.Commit();
            }
            catch (Exception ex)
            {
                Console.WriteLine(ex.Message);
                tx.Rollback();
            }
        }
    }
}
